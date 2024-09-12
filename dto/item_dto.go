package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

/*
UpdateItemInput
updateの場合は、値があれば更新する
そのため、*ポインタ型にすることでnilを許容する

omitnil は型がポインタ型のみ使用できる
値がnilの場合はバリデーションを行わない
*/
type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
