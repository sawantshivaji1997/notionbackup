package node_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/jomei/notionapi"
	"github.com/sawantshivaji1997/notionbackup/src/mocks"
	"github.com/sawantshivaji1997/notionbackup/src/rw"
	"github.com/sawantshivaji1997/notionbackup/src/tree/node"
	"github.com/stretchr/testify/assert"
)

func TestCreateNodeForAllTypes(t *testing.T) {
	tests := []struct {
		name       string
		identifier rw.DataIdentifier
		err        error
		wantErr    bool
	}{
		{
			name:       "Return valid node",
			identifier: rw.DataIdentifier(uuid.NewString()),
			err:        nil,
			wantErr:    false,
		},
		{
			name:       "Return error",
			identifier: "",
			err:        errors.New("error while writing object"),
			wantErr:    true,
		},
	}

	assert := assert.New(t)
	rootNode := node.CreateRootNode()

	assert.NotNil(rootNode)
	assert.Equal(node.NodeID(uuid.Nil.String()), rootNode.GetID())
	assert.Equal(node.NodeType(node.ROOT), rootNode.GetNodeType())
	assert.Empty(rootNode.GetIdentifier())
	assert.False(rootNode.HasChildNode())
	assert.Nil(rootNode.GetChildNode())

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockedRW := mocks.NewReaderWriter(t)

			mockedRW.On("WriteDatabase", context.Background(), &notionapi.Database{}).Return(test.identifier, test.err)
			databaseNode, err1 := node.CreateDatabaseNode(context.Background(), &notionapi.Database{}, mockedRW)

			mockedRW.On("WritePage", context.Background(), &notionapi.Page{}).Return(test.identifier, test.err)
			pageNode, err2 := node.CreatePageNode(context.Background(), &notionapi.Page{}, mockedRW)

			mockedRW.On("WriteBlock", context.Background(), &notionapi.ParagraphBlock{}).Return(test.identifier, test.err)
			blockNode, err3 := node.CreateBlockNode(context.Background(), &notionapi.ParagraphBlock{}, mockedRW)

			mockedRW.AssertExpectations(t)

			if test.wantErr {
				assert.Nil(databaseNode)
				assert.Nil(pageNode)
				assert.Nil(blockNode)
				assert.NotNil(err1)
				assert.NotNil(err2)
				assert.NotNil(err3)
			} else {
				expectedIdentifier := rw.DataIdentifier(test.identifier)
				// Assert DatabaseNode
				assert.NotNil(databaseNode)
				assert.Equal(expectedIdentifier, databaseNode.GetIdentifier())
				assert.Equal(node.NodeType(node.DATABASE), databaseNode.GetNodeType())
				assert.NotEmpty(databaseNode.GetID())
				assert.False(databaseNode.HasChildNode())
				assert.Nil(databaseNode.GetChildNode())
				assert.Nil(err1)

				// Assert PageNode
				assert.NotNil(pageNode)
				assert.Equal(expectedIdentifier, pageNode.GetIdentifier())
				assert.Equal(node.NodeType(node.PAGE), pageNode.GetNodeType())
				assert.NotEmpty(pageNode.GetID())
				assert.False(pageNode.HasChildNode())
				assert.Nil(pageNode.GetChildNode())
				assert.Nil(err2)

				// Assert BlockNode
				assert.NotNil(blockNode)
				assert.Equal(expectedIdentifier, blockNode.GetIdentifier())
				assert.Equal(node.NodeType(node.BLOCK), blockNode.GetNodeType())
				assert.NotEmpty(blockNode.GetID())
				assert.False(blockNode.HasChildNode())
				assert.Nil(blockNode.GetChildNode())
				assert.Nil(err3)
			}
		})
	}
}

func getRandomNodeObject(t *testing.T) *node.Node {
	n := rand.Intn(3)
	mockedRW := mocks.NewReaderWriter(t)

	if n == 0 {
		mockedRW.On("WriteDatabase", context.Background(), &notionapi.Database{}).Return(rw.DataIdentifier(uuid.New().String()), nil)
		databaseNode, _ := node.CreateDatabaseNode(context.Background(), &notionapi.Database{}, mockedRW)
		assert.NotNil(t, databaseNode)
		return databaseNode
	} else if n == 1 {
		mockedRW.On("WritePage", context.Background(), &notionapi.Page{}).Return(rw.DataIdentifier(uuid.New().String()), nil)
		pageNode, _ := node.CreatePageNode(context.Background(), &notionapi.Page{}, mockedRW)
		assert.NotNil(t, pageNode)
		return pageNode
	}
	mockedRW.On("WriteBlock", context.Background(), &notionapi.ParagraphBlock{}).Return(rw.DataIdentifier(uuid.New().String()), nil)
	blockNode, _ := node.CreateBlockNode(context.Background(), &notionapi.ParagraphBlock{}, mockedRW)

	assert.NotNil(t, blockNode)
	return blockNode
}

func TestAddChild(t *testing.T) {
	assert := assert.New(t)
	parentNode := getRandomNodeObject(t)
	child1 := getRandomNodeObject(t)

	parentNode.AddChild(child1)

	assert.False(parentNode.HasSibling())
	assert.True(parentNode.HasChildNode())
	assert.NotNil(parentNode.GetChildNode())
	assert.Equal(parentNode.GetChildNode(), child1)

	child2 := getRandomNodeObject(t)
	parentNode.AddChild(child2)

	assert.True(child1.HasSibling())
	assert.Equal(child1.GetSiblingNode(), child2)

	child3 := getRandomNodeObject(t)
	parentNode.AddChild(child3)
	assert.Equal(child2.GetSiblingNode(), child3)

	n := 1
	temp := parentNode.GetChildNode()
	for temp.HasSibling() {
		n++
		temp = temp.GetSiblingNode()
	}

	assert.Equal(3, n)
}
