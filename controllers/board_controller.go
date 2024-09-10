package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend-onboard-task/dto"
    "backend-onboard-task/helpers"
    "backend-onboard-task/models"
    "backend-onboard-task/services"
)

type BoardController struct {
    BoardService *services.BoardService
}

func NewBoardController(boardService *services.BoardService) *BoardController {
    return &BoardController{
        BoardService: boardService,
    }
}

// 게시글 생성
func (c *BoardController) CreateBoard(ctx *gin.Context) {
    userID, _ := ctx.Get("userID")

    var input dto.CreateBoardInput
    if !helpers.BindJSON(ctx, &input) {
        return
    }

    board := models.Board{
        Title:   input.Title,
        Content: input.Content,
        UserID:  userID.(uint),
    }

    if err := c.BoardService.CreateBoard(board); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "게시글이 성공적으로 생성되었습니다."})
}

// 게시글 목록 조회
func (c *BoardController) GetBoards(ctx *gin.Context) {
    page, limit, err := helpers.GetPaginationParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    boards, total, err := c.BoardService.GetBoards(page, limit)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "boards": boards,
        "total":  total,
        "page":   page,
        "limit":  limit,
        
    })
}

// 게시글 상세 조회
func (c *BoardController) GetBoard(ctx *gin.Context) {
    id := ctx.Param("id")
    board, err := c.BoardService.GetBoardByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"board": board})
}

// 게시글 수정
func (c *BoardController) UpdateBoard(ctx *gin.Context) {
    id := ctx.Param("id")
    userID, _ := ctx.Get("userID")

    var input dto.UpdateBoardInput
    if !helpers.BindJSON(ctx, &input) {
        return
    }

    if input.Title == "" && input.Content == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "제목 또는 본문을 입력해 주세요."})
        return
    }

    if err := c.BoardService.UpdateBoard(id, input, userID.(uint)); err != nil {
        ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "게시글이 성공적으로 수정되었습니다."})
}

// 게시글 삭제
func (c *BoardController) DeleteBoard(ctx *gin.Context) {
    id := ctx.Param("id")
    userID, _ := ctx.Get("userID")

    if err := c.BoardService.DeleteBoard(id, userID.(uint)); err != nil {
        ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "게시글이 성공적으로 삭제되었습니다."})
}