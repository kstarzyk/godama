package wrap

import ui "github.com/gizak/termui"


type TextBox struct {
  box *ui.Par
}

func (tb TextBox) UI() *ui.Par {
  return tb.box
}

func (tb TextBox) Text(newText string) {
  tb.box.BorderLabel = newText
}

func NewTextBox(label, text string, boardless bool) *TextBox {
  tb := &TextBox{}
  tb.box = ui.NewPar(label)
  tb.box.Height = 3
  tb.box.Width = 50
  tb.box.TextFgColor = ui.ColorWhite
  tb.box.BorderLabel = text
  tb.box.BorderFg = ui.ColorCyan
  if boardless {
    tb.box.Border = false
  }
  return tb
}


