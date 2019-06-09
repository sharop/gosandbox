package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/apache/arrow/go/arrow"

	"github.com/apache/arrow/go/arrow/csv"
	"github.com/apache/arrow/go/arrow/memory"
)

func lineCounter(r io.Reader) (int, error) {
	buf :=
		make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err

		}

	}
}

func main() {

	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())

	//raw, err := ioutil.ReadFile("/Users/sharop/Google Drive File Stream/My Drive/Data/SAT_CFDI/tempo.tsv")
	//raw, err := ioutil.ReadFile("/Users/sharop/Google Drive File Stream/My Drive/Data/SAT_CFDI/cddfi_3.csv")
	raw, err := ioutil.ReadFile("/Users/sharop/Google Drive File Stream/My Drive/Data/SAT_CFDI/canasta.json")
	if err != nil {
		log.Fatal(err)
	}
	linesCount, _ := lineCounter(bytes.NewReader(raw))

	fmt.Println("File lines: ", linesCount)

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
			{Name: "rfc_comprador", Type: arrow.BinaryTypes.String},
			{Name: "rfc_vendedor", Type: arrow.BinaryTypes.String},
			{Name: "anos_comprador", Type: arrow.PrimitiveTypes.Int8},
			{Name: "conceptos", Type: arrow.BinaryTypes.String},
			{Name: "edo_comprador", Type: arrow.BinaryTypes.String},
			{Name: "edo_vendedor", Type: arrow.BinaryTypes.String},
			{Name: "fecha_compra", Type: arrow.BinaryTypes.String},
			{Name: "id_anom", Type: arrow.BinaryTypes.String},
			{Name: "inv_distancia", Type: arrow.BinaryTypes.String},
			{Name: "monto_compra", Type: arrow.PrimitiveTypes.Float64},
			{Name: "nombre_comprador", Type: arrow.BinaryTypes.String},
			{Name: "nombre_vendedor", Type: arrow.BinaryTypes.String},
			{Name: "rubro_comprador", Type: arrow.BinaryTypes.String},
			{Name: "rubro_vendedor", Type: arrow.BinaryTypes.String},
			{Name: "tamaño_comprador", Type: arrow.BinaryTypes.String},
			{Name: "tamaño_num_comprador", Type: arrow.BinaryTypes.String},
			{Name: "tamaño_num_vendedor", Type: arrow.BinaryTypes.String},
			{Name: "tamaño_vendedor", Type: arrow.BinaryTypes.String},
			{Name: "version_cfdi", Type: arrow.PrimitiveTypes.Float32},
			{Name: "x_comprador", Type: arrow.PrimitiveTypes.Float32},
			{Name: "x_vendedor", Type: arrow.PrimitiveTypes.Float32},
			{Name: "y_comprador", Type: arrow.PrimitiveTypes.Float32},
			{Name: "y_vendedor", Type: arrow.PrimitiveTypes.Float32},
		},
		nil,
	)

	r := csv.NewReader(
		bytes.NewReader(raw),
		schema,
		csv.WithComma('|'),
		csv.WithHeader(),
		csv.WithChunk(linesCount),
		csv.WithAllocator(mem),
	)

	defer r.Release()
	ni := 0
	for r.Next() {
		rec := r.Record()
		for i, col := range rec.Columns() {
			fmt.Printf("rec[%d][%q]: %v\n", ni, rec.ColumnName(i), col)
		}
		ni++
	}
	//r.Retain()
	//r.Release()

	/*mem := memory.NewGoAllocator()

	schema :=
		arrow.NewSchema(
			[]arrow.Field{
				arrow.Field{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
				arrow.Field{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
			},
			nil,
		)

	b := array.NewRecordBuilder(mem, schema)
	defer b.Release()

	b.Field(0).(*array.Int32Builder).AppendValues(
		[]int32{1, 2, 3, 4, 5, 6},
		nil,
	)

	b.Field(0).(*array.Int32Builder).AppendValues(
		[]int32{7, 8, 9, 10},
		[]bool{true, true, false, true},
	)
	b.Field(1).(*array.Float64Builder).AppendValues(
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		nil,
	)

	rec1 := b.NewRecord()
	defer rec1.Release()

	b.Field(0).(*array.Int32Builder).AppendValues([]int32{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		nil,
	)
	b.Field(1).(*array.Float64Builder).AppendValues([]float64{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		nil,
	)

	rec2 := b.NewRecord()
	defer rec2.Release()

	tbl := array.NewTableFromRecords(schema, []array.Record{rec1, rec2})
	defer tbl.Release()

	tr := array.NewTableReader(tbl, 20)
	defer tr.Release()

	n := 0
	for tr.Next() {
		rec := tr.Record()
		for i, col := range rec.Columns() {
			fmt.Printf("rec[%d][%q]: %v\n", n, rec.ColumnName(i), col)
		}
		n++
	}
	*/
	/*f := bytes.NewBufferString(`## a simple set of data: int64;float64;string
	0;0;str-0
	1;1;str-1
	2;2;str-2
	3;3;str-3
	4;4;str-4
	5;5;str-5
	6;6;str-6
	7;7;str-7
	8;8;str-8
	9;9;str-9
	`)

		schemaCSV := arrow.NewSchema(
			[]arrow.Field{
				{Name: "i64", Type: arrow.PrimitiveTypes.Int64},
				{Name: "f64", Type: arrow.PrimitiveTypes.Float64},
				{Name: "str", Type: arrow.BinaryTypes.String},
			},
			nil, // no metadata
		)

		r := csv.NewReader(
			f, schemaCSV, csv.WithComment('#'), csv.WithComma(';'), csv.WithChunk(10),
		)

		defer r.Release()

		ni := 0
		for r.Next() {
			rec := r.Record()
			for i, col := range rec.Columns() {
				fmt.Printf("rec[%d][%q]: %v\n", ni, rec.ColumnName(i), col)
			}
			ni++
		}*/
}
