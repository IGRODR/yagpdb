// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPremiumSlots(t *testing.T) {
	t.Parallel()

	query := PremiumSlots()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPremiumSlotsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPremiumSlotsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PremiumSlots().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPremiumSlotsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PremiumSlotSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPremiumSlotsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PremiumSlotExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PremiumSlot exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PremiumSlotExists to return true, but got false.")
	}
}

func testPremiumSlotsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	premiumSlotFound, err := FindPremiumSlot(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if premiumSlotFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPremiumSlotsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PremiumSlots().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPremiumSlotsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PremiumSlots().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPremiumSlotsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	premiumSlotOne := &PremiumSlot{}
	premiumSlotTwo := &PremiumSlot{}
	if err = randomize.Struct(seed, premiumSlotOne, premiumSlotDBTypes, false, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}
	if err = randomize.Struct(seed, premiumSlotTwo, premiumSlotDBTypes, false, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = premiumSlotOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = premiumSlotTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PremiumSlots().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPremiumSlotsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	premiumSlotOne := &PremiumSlot{}
	premiumSlotTwo := &PremiumSlot{}
	if err = randomize.Struct(seed, premiumSlotOne, premiumSlotDBTypes, false, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}
	if err = randomize.Struct(seed, premiumSlotTwo, premiumSlotDBTypes, false, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = premiumSlotOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = premiumSlotTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testPremiumSlotsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPremiumSlotsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(premiumSlotColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPremiumSlotToManySlotPremiumCodes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PremiumSlot
	var b, c PremiumCode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, premiumCodeDBTypes, false, premiumCodeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, premiumCodeDBTypes, false, premiumCodeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.SlotID, a.ID)
	queries.Assign(&c.SlotID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SlotPremiumCodes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.SlotID, b.SlotID) {
			bFound = true
		}
		if queries.Equal(v.SlotID, c.SlotID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PremiumSlotSlice{&a}
	if err = a.L.LoadSlotPremiumCodes(ctx, tx, false, (*[]*PremiumSlot)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SlotPremiumCodes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SlotPremiumCodes = nil
	if err = a.L.LoadSlotPremiumCodes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SlotPremiumCodes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPremiumSlotToManyAddOpSlotPremiumCodes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PremiumSlot
	var b, c, d, e PremiumCode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, premiumSlotDBTypes, false, strmangle.SetComplement(premiumSlotPrimaryKeyColumns, premiumSlotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PremiumCode{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, premiumCodeDBTypes, false, strmangle.SetComplement(premiumCodePrimaryKeyColumns, premiumCodeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*PremiumCode{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSlotPremiumCodes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.SlotID) {
			t.Error("foreign key was wrong value", a.ID, first.SlotID)
		}
		if !queries.Equal(a.ID, second.SlotID) {
			t.Error("foreign key was wrong value", a.ID, second.SlotID)
		}

		if first.R.Slot != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Slot != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SlotPremiumCodes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SlotPremiumCodes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SlotPremiumCodes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testPremiumSlotToManySetOpSlotPremiumCodes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PremiumSlot
	var b, c, d, e PremiumCode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, premiumSlotDBTypes, false, strmangle.SetComplement(premiumSlotPrimaryKeyColumns, premiumSlotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PremiumCode{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, premiumCodeDBTypes, false, strmangle.SetComplement(premiumCodePrimaryKeyColumns, premiumCodeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSlotPremiumCodes(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SlotPremiumCodes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSlotPremiumCodes(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SlotPremiumCodes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SlotID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SlotID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.SlotID) {
		t.Error("foreign key was wrong value", a.ID, d.SlotID)
	}
	if !queries.Equal(a.ID, e.SlotID) {
		t.Error("foreign key was wrong value", a.ID, e.SlotID)
	}

	if b.R.Slot != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Slot != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Slot != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Slot != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.SlotPremiumCodes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.SlotPremiumCodes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testPremiumSlotToManyRemoveOpSlotPremiumCodes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PremiumSlot
	var b, c, d, e PremiumCode

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, premiumSlotDBTypes, false, strmangle.SetComplement(premiumSlotPrimaryKeyColumns, premiumSlotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PremiumCode{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, premiumCodeDBTypes, false, strmangle.SetComplement(premiumCodePrimaryKeyColumns, premiumCodeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSlotPremiumCodes(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SlotPremiumCodes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSlotPremiumCodes(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SlotPremiumCodes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SlotID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SlotID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Slot != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Slot != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Slot != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Slot != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.SlotPremiumCodes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.SlotPremiumCodes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.SlotPremiumCodes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testPremiumSlotsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPremiumSlotsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PremiumSlotSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPremiumSlotsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PremiumSlots().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	premiumSlotDBTypes = map[string]string{`ID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `AttachedAt`: `timestamp with time zone`, `UserID`: `bigint`, `GuildID`: `bigint`, `Title`: `text`, `Message`: `text`, `Source`: `text`, `SourceID`: `bigint`, `FullDuration`: `bigint`, `Permanent`: `boolean`, `DurationRemaining`: `bigint`}
	_                  = bytes.MinRead
)

func testPremiumSlotsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(premiumSlotPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(premiumSlotColumns) == len(premiumSlotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPremiumSlotsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(premiumSlotColumns) == len(premiumSlotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PremiumSlot{}
	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, premiumSlotDBTypes, true, premiumSlotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(premiumSlotColumns, premiumSlotPrimaryKeyColumns) {
		fields = premiumSlotColumns
	} else {
		fields = strmangle.SetComplement(
			premiumSlotColumns,
			premiumSlotPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PremiumSlotSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPremiumSlotsUpsert(t *testing.T) {
	t.Parallel()

	if len(premiumSlotColumns) == len(premiumSlotPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PremiumSlot{}
	if err = randomize.Struct(seed, &o, premiumSlotDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PremiumSlot: %s", err)
	}

	count, err := PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, premiumSlotDBTypes, false, premiumSlotPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PremiumSlot struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PremiumSlot: %s", err)
	}

	count, err = PremiumSlots().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
