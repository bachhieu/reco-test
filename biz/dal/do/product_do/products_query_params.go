/* !!
 * Code generated by fkit
 * If you need to write additional functions, should create a separate file.
 * File Created: Sunday, 27 April 2025 16:10:50 +07:00
 * Author: VietLD (leduyviet2612@gmail.com)
 * -----
 * Last Modified: Sunday, 27 April 2025 16:10:50 +07:00
 * Modified By: HIEUBV\hieubv
 * -----
 * Copyright 2025. All rights reserved.
 */

package product_do

type ProductQueryParams struct {
	id       *string `query:"id" json:"id"`
	Q        *string `query:"q" json:"q"`
	Page     int     `query:"page" json:"page"`
	PageSize int     `query:"page_size" json:"page_size"`
}

/* validate */
func (p *ProductQueryParams) Validate() error {
	return nil
}
