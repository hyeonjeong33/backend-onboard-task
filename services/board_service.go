package services

import (
    "gorm.io/gorm"
    "errors"
    "backend-onboard-task/dto"
    "backend-onboard-task/models"
)

type BoardService struct {
    DB *gorm.DB
}

func NewBoardService(db *gorm.DB) *BoardService {
    return &BoardService{
        DB: db,
    }
}

func (s *BoardService) CreateBoard(board models.Board) error {
    if err := s.DB.Create(&board).Error; err != nil {
        return errors.New("게시글 생성에 실패했습니다.")
    }
    return nil
}

func (s *BoardService) GetBoards() ([]dto.BoardResponse, error) {
    var boards []models.Board
    var boardResponses []dto.BoardResponse

    if err := s.DB.Find(&boards).Error; err != nil {
        return nil, errors.New("게시글 목록 조회에 실패했습니다.")
    }

    for _, board := range boards {
        boardResponse := dto.BoardResponse{
            ID:      board.ID,
            Title:   board.Title,
            Content: board.Content,
            Views:   board.Views,
            CreatedAt: board.CreatedAt,
        }
        boardResponses = append(boardResponses, boardResponse)
    }
       
    return boardResponses, nil
}

func (s *BoardService) GetBoardByID(id string) (*dto.BoardResponse, error) {
    var board models.Board
    if err := s.DB.First(&board, id).Error; err != nil {
        return nil, errors.New("해당 게시글을 찾을 수 없습니다.")
    }
    
	board.Views++
	if err := s.DB.Save(&board).Error; err != nil {
		return nil, err
	}

    boardResponse := dto.BoardResponse{
        ID:      board.ID,
        Title:   board.Title,
        Content: board.Content,
        Views:   board.Views,
        CreatedAt: board.CreatedAt,
    }

    return &boardResponse, nil
}

func (s *BoardService) UpdateBoard(id string, input dto.UpdateBoardInput, userID uint) error {
    var existingBoard models.Board
    if err := s.DB.Select("UserID").First(&existingBoard, id).Error; err != nil {
        return errors.New("해당 게시글을 찾을 수 없습니다.")
    }

    if existingBoard.UserID != userID {
        return errors.New("해당 게시글을 수정할 권한이 없습니다.")
    }

    if input.Title != "" {
        existingBoard.Title = input.Title
    }
    if input.Content != "" {
        existingBoard.Content = input.Content
    }

    if err := s.DB.Save(&existingBoard).Error; err != nil {
        return errors.New("게시글 수정에 실패했습니다.")
    }

    return nil
}

func (s *BoardService) DeleteBoard(id string, userID uint) error {
    var board models.Board
    if err := s.DB.Select("UserID").First(&board, id).Error; err != nil {
        return errors.New("해당 게시글을 찾을 수 없습니다.")
    }

    if board.UserID != userID {
        return errors.New("해당 게시글을 삭제할 권한이 없습니다.")
    }

    if err := s.DB.Delete(&board).Error; err != nil {
        return errors.New("게시글 삭제에 실패했습니다.")
    }

    return nil
}