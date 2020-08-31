package miro

import (
	"context"

	"github.com/Miro-Ecosystem/go-miro/miro"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBoard() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Board",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the Board",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
		CreateContext: resourceBoardCreate,
		ReadContext:   resourceBoardRead,
		UpdateContext: resourceBoardUpdate,
		DeleteContext: resourceBoardDelete,
	}
}

func resourceBoardRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	var diags diag.Diagnostics

	board, err := c.Boards.Get(ctx, data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if board == nil {
		data.SetId("")
		return diags
	}

	if err := data.Set("boards", board); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)
	return diags
}

func resourceBoardCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	name := data.Get("name").(string)
	description := data.Get("description").(string)

	req := &miro.CreateBoardRequest{
		Name:        name,
		Description: description,
	}

	board, err := c.Boards.Create(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)
	return resourceBoardRead(ctx, data, meta)
}

func resourceBoardUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	name := data.Get("name").(string)
	description := data.Get("description").(string)

	req := &miro.UpdateBoardRequest{
		Name:        name,
		Description: description,
	}

	_, err := c.Boards.Update(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceBoardRead(ctx, data, meta)
}

func resourceBoardDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)
	var diags diag.Diagnostics
	if err := c.Boards.Delete(ctx, data.Id()); err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")
	return diags
}
