package main

import (
	"errors"
	"time"

	"gofr.dev/pkg/gofr"
)

type Vehicle struct {
	ID             int    `json:"id"`
	Owner          string `json:"owner"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	Vehicle_No     string `json:"vehicle_no"`
	Model          string `json:"model"`
	Defective_Part string `json:"defective_part"`
	Amount         string `json:"amount"`
	Check_in       string `json:"check_in"`
}

type Old_Vehicle struct {
	ID             int    `json:"id"`
	Owner          string `json:"owner"`
	Type           string `json:"type"`
	Color          string `json:"color"`
	Vehicle_No     string `json:"vehicle_no"`
	Model          string `json:"model"`
	Defective_Part string `json:"defective_part"`
	Amount         string `json:"amount"`
	Check_in       string `json:"check_in"`
	Check_out      string `json:"check_out"`
}

func main() {
	app := gofr.New()

	app.POST("/customer/{owner}/{type}/{color}/{vehicle_no}/{model}/{defective_part}/{amount}", func(ctx *gofr.Context) (interface{}, error) {
		owner := ctx.PathParam("owner")
		t := ctx.PathParam("type")
		color := ctx.PathParam("color")
		vehicle_no := ctx.PathParam("vehicle_no")
		model := ctx.PathParam("model")
		defective_part := ctx.PathParam("defective_part")
		amount := ctx.PathParam("amount")
		now := time.Now()

		_, err := ctx.DB().ExecContext(ctx,
			"INSERT INTO garage (owner,type,color,vehicle_no,model,check_in,defective_part,amount) VALUES (?,?,?,?,?,?,?,?)",
			owner, t, color, vehicle_no, model, now.Format("2006-01-02"), defective_part, amount)

		return nil, err
	})

	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var garage []Vehicle

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM garage")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Vehicle
			if err := rows.Scan(&customer.ID, &customer.Owner, &customer.Type, &customer.Color,
				&customer.Vehicle_No, &customer.Model, &customer.Check_in, &customer.Defective_Part,
				&customer.Amount); err != nil {
				return nil, err
			}

			garage = append(garage, customer)
		}

		return garage, nil
	})

	app.GET("/previous", func(ctx *gofr.Context) (interface{}, error) {
		var garage []Old_Vehicle

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM complete")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Old_Vehicle
			if err := rows.Scan(&customer.ID, &customer.Owner, &customer.Type, &customer.Color,
				&customer.Vehicle_No, &customer.Model, &customer.Check_in, &customer.Check_out,
				&customer.Defective_Part, &customer.Amount); err != nil {
				return nil, err
			}

			garage = append(garage, customer)
		}

		return garage, nil
	})

	app.DELETE("/customer/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		now := time.Now()

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM garage WHERE id=?", id)
		if err != nil {
			return nil, err
		}

		var garage []Vehicle

		for rows.Next() {
			var customer Vehicle
			if err := rows.Scan(&customer.ID, &customer.Owner, &customer.Type, &customer.Color,
				&customer.Vehicle_No, &customer.Model, &customer.Check_in, &customer.Defective_Part,
				&customer.Amount); err != nil {
				return nil, err
			}

			garage = append(garage, customer)
		}

		up, err := ctx.DB().ExecContext(ctx,
			"INSERT INTO complete (id,owner,type,color,vehicle_no,model,check_in,check_out,defective_part,amount) VALUES (?,?,?,?,?,?,?,?,?,?)",
			id, garage[0].Owner, garage[0].Type, garage[0].Color, garage[0].Vehicle_No, garage[0].Model,
			garage[0].Check_in, now.Format("2006-01-02"), garage[0].Defective_Part, garage[0].Amount)

		if err != nil {
			return nil, err
		}

		_, err = ctx.DB().ExecContext(ctx,
			"DELETE FROM garage WHERE id=?",
			id)

		if err != nil {
			return nil, err
		}

		return up, nil
	})

	app.PUT("/customer/{id}/{item}/{new_item}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		item := ctx.PathParam("item")
		new_item := ctx.PathParam("new_item")

		rows, err := ctx.DB().QueryContext(ctx, "SELECT owner,type,color,vehicle_no,model,defective_part,amount FROM garage WHERE id=?", id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var garage []Vehicle

		for rows.Next() {
			var customer Vehicle
			if err := rows.Scan(&customer.Owner, &customer.Type, &customer.Color,
				&customer.Vehicle_No, &customer.Model, &customer.Defective_Part,
				&customer.Amount); err != nil {
				return nil, err
			}

			garage = append(garage, customer)
		}

		if len(garage) == 0 {
			return nil, errors.New("Record not found for the given ID")
		}

		owner := garage[0].Owner
		t := garage[0].Type
		color := garage[0].Color
		vehicle_no := garage[0].Vehicle_No
		model := garage[0].Model
		defective_part := garage[0].Defective_Part
		amount := garage[0].Amount

		switch item {
		case "owner":
			owner = new_item
		case "type":
			t = new_item
		case "color":
			color = new_item
		case "vehicle_no":
			vehicle_no = new_item
		case "model":
			model = new_item
		case "defective_part":
			defective_part = new_item
		case "amount":
			amount = new_item
		}

		_, err = ctx.DB().ExecContext(ctx,
			"UPDATE garage SET owner=?, type=?, color=?, vehicle_no=?, model=?, defective_part=?, amount=? WHERE id=?",
			owner, t, color, vehicle_no, model, defective_part, amount, id)

		return nil, err
	})

	app.Start()
}
