package wrap
import ui "github.com/gizak/termui"

type CList struct {
  list *ui.List
  selected int
  isFocused bool
}

func (c *CList) UI() *ui.List {
  return c.list
}

func NewCList(label string, items []string) *CList {
  l := &CList{}

  l.selected = 0
  l.isFocused = false
  l.list = ui.NewList()
  l.list.Items = items
  l.list.ItemFgColor = ui.ColorYellow
  l.list.BorderLabel = label
  l.list.Height = 7
  l.list.Width = 25
  l.list.Y = 0
  return l
}

func (c* CList) Focus() error {
 c.isFocused = true
 return nil
}
func (c* CList) Defocus() error {
  c.isFocused = false
  c.selected = 0
  return nil
}
    //tables.list.Items[cnt] = "["+tables.list.Items[cnt]+"]" + "(fg-blue)"

func (c *CList) HandleKey(key string) error {
  switch key {
  case "k": c.list.Width += 1
  case "j": c.list.Width -= 1
  case "<up>":
    if c.selected != 0 {
      c.selected--;
    } else {
      c.selected = len(c.list.Items)-1
    }
  case "<down>":
    if c.selected != len(c.list.Items)-1 {
      c.selected++;
    } else {
      c.selected = 0
    }
  }
  return nil
}
