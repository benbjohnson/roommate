package roommate
import (
"fmt"
"io"
)
//line error.ego:1
 func (t *Template) Error(w io.Writer, err error) error  {
//line error.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line error.ego:3
if _, err := fmt.Fprintf(w, "<html>\n"); err != nil { return err }
//line error.ego:4
 t.Head(w) 
//line error.ego:5
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line error.ego:6
if _, err := fmt.Fprintf(w, "<body>\n  "); err != nil { return err }
//line error.ego:7
if _, err := fmt.Fprintf(w, "<h1>Roommate"); err != nil { return err }
//line error.ego:7
if _, err := fmt.Fprintf(w, "</h1>\n\n  "); err != nil { return err }
//line error.ego:9
if _, err := fmt.Fprintf(w, "<h2>An Error Occurred"); err != nil { return err }
//line error.ego:9
if _, err := fmt.Fprintf(w, "</h2>\n\n  "); err != nil { return err }
//line error.ego:11
if _, err := fmt.Fprintf(w, "<p>"); err != nil { return err }
//line error.ego:11
if _, err := fmt.Fprintf(w, "%v",  err ); err != nil { return err }
//line error.ego:11
if _, err := fmt.Fprintf(w, "</p>\n"); err != nil { return err }
//line error.ego:12
if _, err := fmt.Fprintf(w, "</body>\n"); err != nil { return err }
//line error.ego:13
if _, err := fmt.Fprintf(w, "</html>\n"); err != nil { return err }
return nil
}
//line head.ego:1
 func (t *Template) Head(w io.Writer) error  {
//line head.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line head.ego:3
if _, err := fmt.Fprintf(w, "<head>\n"); err != nil { return err }
//line head.ego:4
if _, err := fmt.Fprintf(w, "<title>Roommate"); err != nil { return err }
//line head.ego:4
if _, err := fmt.Fprintf(w, "</title>\n"); err != nil { return err }
//line head.ego:5
if _, err := fmt.Fprintf(w, "</head>\n"); err != nil { return err }
return nil
}
//line index.ego:1
 func (t *Template) Index(w io.Writer, rooms []*Room) error  {
//line index.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line index.ego:3
if _, err := fmt.Fprintf(w, "<html>\n"); err != nil { return err }
//line index.ego:4
 t.Head(w) 
//line index.ego:5
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line index.ego:6
if _, err := fmt.Fprintf(w, "<body>\n  "); err != nil { return err }
//line index.ego:7
if _, err := fmt.Fprintf(w, "<h1>Roommate"); err != nil { return err }
//line index.ego:7
if _, err := fmt.Fprintf(w, "</h1>\n\n  "); err != nil { return err }
//line index.ego:9
if _, err := fmt.Fprintf(w, "<h2>Rooms"); err != nil { return err }
//line index.ego:9
if _, err := fmt.Fprintf(w, "</h2>\n  "); err != nil { return err }
//line index.ego:10
 for _, r := range rooms { 
//line index.ego:11
if _, err := fmt.Fprintf(w, "\n    "); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "<li>"); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "<a href=\"/rooms/"); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "%v",  r.ID ); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "\">"); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "%v",  r.Name ); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "</a>"); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "</li>\n  "); err != nil { return err }
//line index.ego:12
 } 
//line index.ego:13
if _, err := fmt.Fprintf(w, "\n"); err != nil { return err }
//line index.ego:13
if _, err := fmt.Fprintf(w, "</body>\n"); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "</html>\n"); err != nil { return err }
return nil
}
//line edit.ego:1
 func (t *RoomTemplate) Edit(w io.Writer, r *Room) error  {
//line edit.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line edit.ego:3
if _, err := fmt.Fprintf(w, "<html>\n"); err != nil { return err }
//line edit.ego:4
 t.Head(w) 
//line edit.ego:5
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line edit.ego:6
if _, err := fmt.Fprintf(w, "<body>\n  "); err != nil { return err }
//line edit.ego:7
if _, err := fmt.Fprintf(w, "<h1>Roommate"); err != nil { return err }
//line edit.ego:7
if _, err := fmt.Fprintf(w, "</h1>\n\n  "); err != nil { return err }
//line edit.ego:9
if _, err := fmt.Fprintf(w, "<h2>Edit Room"); err != nil { return err }
//line edit.ego:9
if _, err := fmt.Fprintf(w, "</h2>\n\n  "); err != nil { return err }
//line edit.ego:11
if _, err := fmt.Fprintf(w, "<form method=\"POST\" action=\"/rooms/"); err != nil { return err }
//line edit.ego:11
if _, err := fmt.Fprintf(w, "%v",  r.ID ); err != nil { return err }
//line edit.ego:11
if _, err := fmt.Fprintf(w, "\">\n    "); err != nil { return err }
//line edit.ego:12
if _, err := fmt.Fprintf(w, "<p>\n      "); err != nil { return err }
//line edit.ego:13
if _, err := fmt.Fprintf(w, "<label>Name"); err != nil { return err }
//line edit.ego:13
if _, err := fmt.Fprintf(w, "</label>\n      "); err != nil { return err }
//line edit.ego:14
if _, err := fmt.Fprintf(w, "<input type=\"text\" name=\"name\" value=\""); err != nil { return err }
//line edit.ego:14
if _, err := fmt.Fprintf(w, "%v",  r.Name ); err != nil { return err }
//line edit.ego:14
if _, err := fmt.Fprintf(w, "\"/>\n    "); err != nil { return err }
//line edit.ego:15
if _, err := fmt.Fprintf(w, "</p>\n    "); err != nil { return err }
//line edit.ego:16
if _, err := fmt.Fprintf(w, "<p>\n      "); err != nil { return err }
//line edit.ego:17
if _, err := fmt.Fprintf(w, "<input type=\"submit\" value=\"Update Room\"/>\n    "); err != nil { return err }
//line edit.ego:18
if _, err := fmt.Fprintf(w, "</p>\n  "); err != nil { return err }
//line edit.ego:19
if _, err := fmt.Fprintf(w, "</form>\n"); err != nil { return err }
//line edit.ego:20
if _, err := fmt.Fprintf(w, "</body>\n"); err != nil { return err }
//line edit.ego:21
if _, err := fmt.Fprintf(w, "</html>\n"); err != nil { return err }
return nil
}
//line index.ego:1
 func (t *RoomTemplate) Index(w io.Writer, rooms []*Room) error  {
//line index.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line index.ego:3
if _, err := fmt.Fprintf(w, "<html>\n"); err != nil { return err }
//line index.ego:4
 t.Head(w) 
//line index.ego:5
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line index.ego:6
if _, err := fmt.Fprintf(w, "<body>\n  "); err != nil { return err }
//line index.ego:7
if _, err := fmt.Fprintf(w, "<h1>Roommate"); err != nil { return err }
//line index.ego:7
if _, err := fmt.Fprintf(w, "</h1>\n\n  "); err != nil { return err }
//line index.ego:9
if _, err := fmt.Fprintf(w, "<h2>Room Listing"); err != nil { return err }
//line index.ego:9
if _, err := fmt.Fprintf(w, "</h2>\n  "); err != nil { return err }
//line index.ego:10
 if len(rooms) == 0 { 
//line index.ego:11
if _, err := fmt.Fprintf(w, "\n    "); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "<p>No rooms found."); err != nil { return err }
//line index.ego:11
if _, err := fmt.Fprintf(w, "</p>\n  "); err != nil { return err }
//line index.ego:12
 } else { 
//line index.ego:13
if _, err := fmt.Fprintf(w, "\n    "); err != nil { return err }
//line index.ego:13
 for _, r := range rooms { 
//line index.ego:14
if _, err := fmt.Fprintf(w, "\n      "); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "<li>"); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "<a href=\"/rooms/"); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "%v",  r.ID ); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "\">"); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "%v",  r.Name ); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "</a>"); err != nil { return err }
//line index.ego:14
if _, err := fmt.Fprintf(w, "</li>\n    "); err != nil { return err }
//line index.ego:15
 } 
//line index.ego:16
if _, err := fmt.Fprintf(w, "\n  "); err != nil { return err }
//line index.ego:16
 } 
//line index.ego:17
if _, err := fmt.Fprintf(w, "\n\n  "); err != nil { return err }
//line index.ego:18
if _, err := fmt.Fprintf(w, "<p>\n    "); err != nil { return err }
//line index.ego:19
if _, err := fmt.Fprintf(w, "<a href=\"/rooms/new\">Add a room"); err != nil { return err }
//line index.ego:19
if _, err := fmt.Fprintf(w, "</a>\n  "); err != nil { return err }
//line index.ego:20
if _, err := fmt.Fprintf(w, "</p>\n"); err != nil { return err }
//line index.ego:21
if _, err := fmt.Fprintf(w, "</body>\n"); err != nil { return err }
//line index.ego:22
if _, err := fmt.Fprintf(w, "</html>\n"); err != nil { return err }
return nil
}
//line new.ego:1
 func (t *RoomTemplate) New(w io.Writer) error  {
//line new.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line new.ego:3
if _, err := fmt.Fprintf(w, "<html>\n"); err != nil { return err }
//line new.ego:4
 t.Head(w) 
//line new.ego:5
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line new.ego:6
if _, err := fmt.Fprintf(w, "<body>\n  "); err != nil { return err }
//line new.ego:7
if _, err := fmt.Fprintf(w, "<h1>Roommate"); err != nil { return err }
//line new.ego:7
if _, err := fmt.Fprintf(w, "</h1>\n\n  "); err != nil { return err }
//line new.ego:9
if _, err := fmt.Fprintf(w, "<h2>New Room"); err != nil { return err }
//line new.ego:9
if _, err := fmt.Fprintf(w, "</h2>\n\n  "); err != nil { return err }
//line new.ego:11
if _, err := fmt.Fprintf(w, "<form method=\"POST\" action=\"/rooms\">\n    "); err != nil { return err }
//line new.ego:12
if _, err := fmt.Fprintf(w, "<p>\n      "); err != nil { return err }
//line new.ego:13
if _, err := fmt.Fprintf(w, "<label>Name"); err != nil { return err }
//line new.ego:13
if _, err := fmt.Fprintf(w, "</label>\n      "); err != nil { return err }
//line new.ego:14
if _, err := fmt.Fprintf(w, "<input type=\"text\" name=\"name\"/>\n    "); err != nil { return err }
//line new.ego:15
if _, err := fmt.Fprintf(w, "</p>\n    "); err != nil { return err }
//line new.ego:16
if _, err := fmt.Fprintf(w, "<p>\n      "); err != nil { return err }
//line new.ego:17
if _, err := fmt.Fprintf(w, "<input type=\"submit\" value=\"Create Room\"/>\n    "); err != nil { return err }
//line new.ego:18
if _, err := fmt.Fprintf(w, "</p>\n  "); err != nil { return err }
//line new.ego:19
if _, err := fmt.Fprintf(w, "</form>\n"); err != nil { return err }
//line new.ego:20
if _, err := fmt.Fprintf(w, "</body>\n"); err != nil { return err }
//line new.ego:21
if _, err := fmt.Fprintf(w, "</html>\n"); err != nil { return err }
return nil
}
//line show.ego:1
 func (t *RoomTemplate) Show(w io.Writer, r *Room) error  {
//line show.ego:2
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line show.ego:3
if _, err := fmt.Fprintf(w, "<html>\n"); err != nil { return err }
//line show.ego:4
 t.Head(w) 
//line show.ego:5
if _, err := fmt.Fprintf(w, "\n\n"); err != nil { return err }
//line show.ego:6
if _, err := fmt.Fprintf(w, "<body>\n  "); err != nil { return err }
//line show.ego:7
if _, err := fmt.Fprintf(w, "<h1>Roommate"); err != nil { return err }
//line show.ego:7
if _, err := fmt.Fprintf(w, "</h1>\n\n  "); err != nil { return err }
//line show.ego:9
if _, err := fmt.Fprintf(w, "<h2>"); err != nil { return err }
//line show.ego:9
if _, err := fmt.Fprintf(w, "%v",  r.Name ); err != nil { return err }
//line show.ego:9
if _, err := fmt.Fprintf(w, "</h2>\n\n  "); err != nil { return err }
//line show.ego:11
if _, err := fmt.Fprintf(w, "<p>\n    "); err != nil { return err }
//line show.ego:12
if _, err := fmt.Fprintf(w, "<a href=\"/rooms\">View all rooms"); err != nil { return err }
//line show.ego:12
if _, err := fmt.Fprintf(w, "</a> |\n    "); err != nil { return err }
//line show.ego:13
if _, err := fmt.Fprintf(w, "<a href=\"/rooms/"); err != nil { return err }
//line show.ego:13
if _, err := fmt.Fprintf(w, "%v",  r.ID ); err != nil { return err }
//line show.ego:13
if _, err := fmt.Fprintf(w, "/edit\">Edit room"); err != nil { return err }
//line show.ego:13
if _, err := fmt.Fprintf(w, "</a> |\n    "); err != nil { return err }
//line show.ego:14
if _, err := fmt.Fprintf(w, "<a href=\"/rooms/"); err != nil { return err }
//line show.ego:14
if _, err := fmt.Fprintf(w, "%v",  r.ID ); err != nil { return err }
//line show.ego:14
if _, err := fmt.Fprintf(w, "/delete\">Delete room"); err != nil { return err }
//line show.ego:14
if _, err := fmt.Fprintf(w, "</a>\n  "); err != nil { return err }
//line show.ego:15
if _, err := fmt.Fprintf(w, "</p>\n"); err != nil { return err }
//line show.ego:16
if _, err := fmt.Fprintf(w, "</body>\n"); err != nil { return err }
//line show.ego:17
if _, err := fmt.Fprintf(w, "</html>\n"); err != nil { return err }
return nil
}
