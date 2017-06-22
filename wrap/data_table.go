package wrap

import ui "github.com/gizak/termui"

type DataTable struct {
  Table *ui.Table
}

func NewDataTable() *DataTable {

  dt := &DataTable{}
  dt.Table = ui.NewTable()
  dt.Table.FgColor = ui.ColorWhite
  dt.Table.BgColor = ui.ColorDefault

  return dt
}


func (dt *DataTable) SetHeader(headers []string) {
  var newList [][]string
  newList = append(newList, headers)
  dt.Table.Rows = newList
}

func (dt *DataTable) SetData(data [][]string) {
  for _, str := range data {
    dt.Table.Rows = append(dt.Table.Rows, str)
  }
}

func (dt *DataTable) SetPosition(X,Y int) {
  dt.Table.Y = Y
  dt.Table.X = X
}

func (dt *DataTable) SetSize(Width,Height int) {
  dt.Table.Width = Width
  dt.Table.Height = Height
}
