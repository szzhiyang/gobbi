# `atk` enums

## `CoordType`

Specifies how xy coordinates are to be interpreted. Used by functions such
as atk_component_get_position() and atk_text_get_character_extents()

C - `AtkCoordType`

### ATK_XY_SCREEN = 0
specifies xy coordinates relative to the screen

### ATK_XY_WINDOW = 1
specifies xy coordinates relative to the widget's
top-level window


## `KeyEventType`

Specifies the type of a keyboard evemt.

C - `AtkKeyEventType`

### ATK_KEY_EVENT_PRESS = 0
specifies a key press event

### ATK_KEY_EVENT_RELEASE = 1
specifies a key release event

### ATK_KEY_EVENT_LAST_DEFINED = 2
Not a valid value; specifies end of enumeration


## `Layer`

Describes the layer of a component

These enumerated "layer values" are used when determining which UI
rendering layer a component is drawn into, which can help in making
determinations of when components occlude one another.

C - `AtkLayer`

### ATK_LAYER_INVALID = 0
The object does not have a layer

### ATK_LAYER_BACKGROUND = 1
This layer is reserved for the desktop background

### ATK_LAYER_CANVAS = 2
This layer is used for Canvas components

### ATK_LAYER_WIDGET = 3
This layer is normally used for components

### ATK_LAYER_MDI = 4
This layer is used for layered components

### ATK_LAYER_POPUP = 5
This layer is used for popup components, such as menus

### ATK_LAYER_OVERLAY = 6
This layer is reserved for future use.

### ATK_LAYER_WINDOW = 7
This layer is used for toplevel windows.


## `RelationType`

Describes the type of the relation

C - `AtkRelationType`

### ATK_RELATION_NULL = 0
Not used, represens "no relationship" or an error condition.

### ATK_RELATION_CONTROLLED_BY = 1
Indicates an object controlled by one or more target objects.

### ATK_RELATION_CONTROLLER_FOR = 2
Indicates an object is an controller for one or more target objects.

### ATK_RELATION_LABEL_FOR = 3
Indicates an object is a label for one or more target objects.

### ATK_RELATION_LABELLED_BY = 4
Indicates an object is labelled by one or more target objects.

### ATK_RELATION_MEMBER_OF = 5
Indicates an object is a member of a group of one or more target objects.

### ATK_RELATION_NODE_CHILD_OF = 6
Indicates an object is a cell in a treetable which is displayed because a cell in the same column is expanded and identifies that cell.

### ATK_RELATION_FLOWS_TO = 7
Indicates that the object has content that flows logically to another
 AtkObject in a sequential way, (for instance text-flow).

### ATK_RELATION_FLOWS_FROM = 8
Indicates that the object has content that flows logically from
 another AtkObject in a sequential way, (for instance text-flow).

### ATK_RELATION_SUBWINDOW_OF = 9
Indicates a subwindow attached to a component but otherwise has no connection in  the UI heirarchy to that component.

### ATK_RELATION_EMBEDS = 10
Indicates that the object visually embeds
 another object's content, i.e. this object's content flows around
 another's content.

### ATK_RELATION_EMBEDDED_BY = 11
Reciprocal of %ATK_RELATION_EMBEDS, indicates that
 this object's content is visualy embedded in another object.

### ATK_RELATION_POPUP_FOR = 12
Indicates that an object is a popup for another object.

### ATK_RELATION_PARENT_WINDOW_OF = 13
Indicates that an object is a parent window of another object.

### ATK_RELATION_DESCRIBED_BY = 14
Reciprocal of %ATK_RELATION_DESCRIPTION_FOR. Indicates that one
or more target objects provide descriptive information about this object. This relation
type is most appropriate for information that is not essential as its presentation may
be user-configurable and/or limited to an on-demand mechanism such as an assistive
technology command. For brief, essential information such as can be found in a widget's
on-screen label, use %ATK_RELATION_LABELLED_BY. For an on-screen error message, use
%ATK_RELATION_ERROR_MESSAGE. For lengthy extended descriptive information contained in
an on-screen object, consider using %ATK_RELATION_DETAILS as assistive technologies may
provide a means for the user to navigate to objects containing detailed descriptions so
that their content can be more closely reviewed.

### ATK_RELATION_DESCRIPTION_FOR = 15
Reciprocal of %ATK_RELATION_DESCRIBED_BY. Indicates that this
object provides descriptive information about the target object(s). See also
%ATK_RELATION_DETAILS_FOR and %ATK_RELATION_ERROR_FOR.

### ATK_RELATION_NODE_PARENT_OF = 16
Indicates an object is a cell in a treetable and is expanded to display other cells in the same column.

### ATK_RELATION_DETAILS = 17
Reciprocal of %ATK_RELATION_DETAILS_FOR. Indicates that this object
has a detailed or extended description, the contents of which can be found in the target
object(s). This relation type is most appropriate for information that is sufficiently
lengthy as to make navigation to the container of that information desirable. For less
verbose information suitable for announcement only, see %ATK_RELATION_DESCRIBED_BY. If
the detailed information describes an error condition, %ATK_RELATION_ERROR_FOR should be
used instead. @Since: ATK-2.26.

### ATK_RELATION_DETAILS_FOR = 18
Reciprocal of %ATK_RELATION_DETAILS. Indicates that this object
provides a detailed or extended description about the target object(s). See also
%ATK_RELATION_DESCRIPTION_FOR and %ATK_RELATION_ERROR_FOR. @Since: ATK-2.26.

### ATK_RELATION_ERROR_MESSAGE = 19
Reciprocal of %ATK_RELATION_ERROR_FOR. Indicates that this object
has one or more errors, the nature of which is described in the contents of the target
object(s). Objects that have this relation type should also contain %ATK_STATE_INVALID_ENTRY
in their #AtkStateSet. @Since: ATK-2.26.

### ATK_RELATION_ERROR_FOR = 20
Reciprocal of %ATK_RELATION_ERROR_MESSAGE. Indicates that this object
contains an error message describing an invalid condition in the target object(s). @Since:
ATK_2.26.

### ATK_RELATION_LAST_DEFINED = 21
Not used, this value indicates the end of the enumeration.


## `Role`

Describes the role of an object

These are the built-in enumerated roles that UI components can have in
ATK.  Other roles may be added at runtime, so an AtkRole >=
ATK_ROLE_LAST_DEFINED is not necessarily an error.

C - `AtkRole`

### ATK_ROLE_INVALID = 0
Invalid role

### ATK_ROLE_ACCEL_LABEL = 1
A label which represents an accelerator

### ATK_ROLE_ALERT = 2
An object which is an alert to the user. Assistive Technologies typically respond to ATK_ROLE_ALERT by reading the entire onscreen contents of containers advertising this role.  Should be used for warning dialogs, etc.

### ATK_ROLE_ANIMATION = 3
An object which is an animated image

### ATK_ROLE_ARROW = 4
An arrow in one of the four cardinal directions

### ATK_ROLE_CALENDAR = 5
An object that displays a calendar and allows the user to select a date

### ATK_ROLE_CANVAS = 6
An object that can be drawn into and is used to trap events

### ATK_ROLE_CHECK_BOX = 7
A choice that can be checked or unchecked and provides a separate indicator for the current state

### ATK_ROLE_CHECK_MENU_ITEM = 8
A menu item with a check box

### ATK_ROLE_COLOR_CHOOSER = 9
A specialized dialog that lets the user choose a color

### ATK_ROLE_COLUMN_HEADER = 10
The header for a column of data

### ATK_ROLE_COMBO_BOX = 11
A collapsible list of choices the user can select from

### ATK_ROLE_DATE_EDITOR = 12
An object whose purpose is to allow a user to edit a date

### ATK_ROLE_DESKTOP_ICON = 13
An inconifed internal frame within a DESKTOP_PANE

### ATK_ROLE_DESKTOP_FRAME = 14
A pane that supports internal frames and iconified versions of those internal frames

### ATK_ROLE_DIAL = 15
An object whose purpose is to allow a user to set a value

### ATK_ROLE_DIALOG = 16
A top level window with title bar and a border

### ATK_ROLE_DIRECTORY_PANE = 17
A pane that allows the user to navigate through and select the contents of a directory

### ATK_ROLE_DRAWING_AREA = 18
An object used for drawing custom user interface elements

### ATK_ROLE_FILE_CHOOSER = 19
A specialized dialog that lets the user choose a file

### ATK_ROLE_FILLER = 20
A object that fills up space in a user interface

### ATK_ROLE_FONT_CHOOSER = 21
A specialized dialog that lets the user choose a font

### ATK_ROLE_FRAME = 22
A top level window with a title bar, border, menubar, etc.

### ATK_ROLE_GLASS_PANE = 23
A pane that is guaranteed to be painted on top of all panes beneath it

### ATK_ROLE_HTML_CONTAINER = 24
A document container for HTML, whose children represent the document content

### ATK_ROLE_ICON = 25
A small fixed size picture, typically used to decorate components

### ATK_ROLE_IMAGE = 26
An object whose primary purpose is to display an image

### ATK_ROLE_INTERNAL_FRAME = 27
A frame-like object that is clipped by a desktop pane

### ATK_ROLE_LABEL = 28
An object used to present an icon or short string in an interface

### ATK_ROLE_LAYERED_PANE = 29
A specialized pane that allows its children to be drawn in layers, providing a form of stacking order

### ATK_ROLE_LIST = 30
An object that presents a list of objects to the user and allows the user to select one or more of them

### ATK_ROLE_LIST_ITEM = 31
An object that represents an element of a list

### ATK_ROLE_MENU = 32
An object usually found inside a menu bar that contains a list of actions the user can choose from

### ATK_ROLE_MENU_BAR = 33
An object usually drawn at the top of the primary dialog box of an application that contains a list of menus the user can choose from

### ATK_ROLE_MENU_ITEM = 34
An object usually contained in a menu that presents an action the user can choose

### ATK_ROLE_OPTION_PANE = 35
A specialized pane whose primary use is inside a DIALOG

### ATK_ROLE_PAGE_TAB = 36
An object that is a child of a page tab list

### ATK_ROLE_PAGE_TAB_LIST = 37
An object that presents a series of panels (or page tabs), one at a time, through some mechanism provided by the object

### ATK_ROLE_PANEL = 38
A generic container that is often used to group objects

### ATK_ROLE_PASSWORD_TEXT = 39
A text object uses for passwords, or other places where the text content is not shown visibly to the user

### ATK_ROLE_POPUP_MENU = 40
A temporary window that is usually used to offer the user a list of choices, and then hides when the user selects one of those choices

### ATK_ROLE_PROGRESS_BAR = 41
An object used to indicate how much of a task has been completed

### ATK_ROLE_PUSH_BUTTON = 42
An object the user can manipulate to tell the application to do something

### ATK_ROLE_RADIO_BUTTON = 43
A specialized check box that will cause other radio buttons in the same group to become unchecked when this one is checked

### ATK_ROLE_RADIO_MENU_ITEM = 44
A check menu item which belongs to a group. At each instant exactly one of the radio menu items from a group is selected

### ATK_ROLE_ROOT_PANE = 45
A specialized pane that has a glass pane and a layered pane as its children

### ATK_ROLE_ROW_HEADER = 46
The header for a row of data

### ATK_ROLE_SCROLL_BAR = 47
An object usually used to allow a user to incrementally view a large amount of data.

### ATK_ROLE_SCROLL_PANE = 48
An object that allows a user to incrementally view a large amount of information

### ATK_ROLE_SEPARATOR = 49
An object usually contained in a menu to provide a visible and logical separation of the contents in a menu

### ATK_ROLE_SLIDER = 50
An object that allows the user to select from a bounded range

### ATK_ROLE_SPLIT_PANE = 51
A specialized panel that presents two other panels at the same time

### ATK_ROLE_SPIN_BUTTON = 52
An object used to get an integer or floating point number from the user

### ATK_ROLE_STATUSBAR = 53
An object which reports messages of minor importance to the user

### ATK_ROLE_TABLE = 54
An object used to represent information in terms of rows and columns

### ATK_ROLE_TABLE_CELL = 55
A cell in a table

### ATK_ROLE_TABLE_COLUMN_HEADER = 56
The header for a column of a table

### ATK_ROLE_TABLE_ROW_HEADER = 57
The header for a row of a table

### ATK_ROLE_TEAR_OFF_MENU_ITEM = 58
A menu item used to tear off and reattach its menu

### ATK_ROLE_TERMINAL = 59
An object that represents an accessible terminal.  @Since: ATK-0.6

### ATK_ROLE_TEXT = 60
An interactive widget that supports multiple lines of text and
optionally accepts user input, but whose purpose is not to solicit user input.
Thus ATK_ROLE_TEXT is appropriate for the text view in a plain text editor
but inappropriate for an input field in a dialog box or web form. For widgets
whose purpose is to solicit input from the user, see ATK_ROLE_ENTRY and
ATK_ROLE_PASSWORD_TEXT. For generic objects which display a brief amount of
textual information, see ATK_ROLE_STATIC.

### ATK_ROLE_TOGGLE_BUTTON = 61
A specialized push button that can be checked or unchecked, but does not provide a separate indicator for the current state

### ATK_ROLE_TOOL_BAR = 62
A bar or palette usually composed of push buttons or toggle buttons

### ATK_ROLE_TOOL_TIP = 63
An object that provides information about another object

### ATK_ROLE_TREE = 64
An object used to represent hierarchical information to the user

### ATK_ROLE_TREE_TABLE = 65
An object capable of expanding and collapsing rows as well as showing multiple columns of data.   @Since: ATK-0.7

### ATK_ROLE_UNKNOWN = 66
The object contains some Accessible information, but its role is not known

### ATK_ROLE_VIEWPORT = 67
An object usually used in a scroll pane

### ATK_ROLE_WINDOW = 68
A top level window with no title or border.

### ATK_ROLE_HEADER = 69
An object that serves as a document header. @Since: ATK-1.1.1

### ATK_ROLE_FOOTER = 70
An object that serves as a document footer.  @Since: ATK-1.1.1

### ATK_ROLE_PARAGRAPH = 71
An object which is contains a paragraph of text content.   @Since: ATK-1.1.1

### ATK_ROLE_RULER = 72
An object which describes margins and tab stops, etc. for text objects which it controls (should have CONTROLLER_FOR relation to such).   @Since: ATK-1.1.1

### ATK_ROLE_APPLICATION = 73
The object is an application object, which may contain @ATK_ROLE_FRAME objects or other types of accessibles.  The root accessible of any application's ATK hierarchy should have ATK_ROLE_APPLICATION.   @Since: ATK-1.1.4

### ATK_ROLE_AUTOCOMPLETE = 74
The object is a dialog or list containing items for insertion into an entry widget, for instance a list of words for completion of a text entry.   @Since: ATK-1.3

### ATK_ROLE_EDITBAR = 75
The object is an editable text object in a toolbar.  @Since: ATK-1.5

### ATK_ROLE_EMBEDDED = 76
The object is an embedded container within a document or panel.  This role is a grouping "hint" indicating that the contained objects share a context.  @Since: ATK-1.7.2

### ATK_ROLE_ENTRY = 77
The object is a component whose textual content may be entered or modified by the user, provided @ATK_STATE_EDITABLE is present.   @Since: ATK-1.11

### ATK_ROLE_CHART = 78
The object is a graphical depiction of quantitative data. It may contain multiple subelements whose attributes and/or description may be queried to obtain both the quantitative data and information about how the data is being presented. The LABELLED_BY relation is particularly important in interpreting objects of this type, as is the accessible-description property.  @Since: ATK-1.11

### ATK_ROLE_CAPTION = 79
The object contains descriptive information, usually textual, about another user interface element such as a table, chart, or image.  @Since: ATK-1.11

### ATK_ROLE_DOCUMENT_FRAME = 80
The object is a visual frame or container which contains a view of document content. Document frames may occur within another Document instance, in which case the second document may be said to be embedded in the containing instance. HTML frames are often ROLE_DOCUMENT_FRAME. Either this object, or a singleton descendant, should implement the Document interface.  @Since: ATK-1.11

### ATK_ROLE_HEADING = 81
The object serves as a heading for content which follows it in a document. The 'heading level' of the heading, if availabe, may be obtained by querying the object's attributes.

### ATK_ROLE_PAGE = 82
The object is a containing instance which encapsulates a page of information. @ATK_ROLE_PAGE is used in documents and content which support a paginated navigation model.  @Since: ATK-1.11

### ATK_ROLE_SECTION = 83
The object is a containing instance of document content which constitutes a particular 'logical' section of the document. The type of content within a section, and the nature of the section division itself, may be obtained by querying the object's attributes. Sections may be nested. @Since: ATK-1.11

### ATK_ROLE_REDUNDANT_OBJECT = 84
The object is redundant with another object in the hierarchy, and is exposed for purely technical reasons.  Objects of this role should normally be ignored by clients. @Since: ATK-1.11

### ATK_ROLE_FORM = 85
The object is a container for form controls, for instance as part of a
web form or user-input form within a document.  This role is primarily a tag/convenience for
clients when navigating complex documents, it is not expected that ordinary GUI containers will
always have ATK_ROLE_FORM. @Since: ATK-1.12.0

### ATK_ROLE_LINK = 86
The object is a hypertext anchor, i.e. a "link" in a
hypertext document.  Such objects are distinct from 'inline'
content which may also use the Hypertext/Hyperlink interfaces
to indicate the range/location within a text object where
an inline or embedded object lies.  @Since: ATK-1.12.1

### ATK_ROLE_INPUT_METHOD_WINDOW = 87
The object is a window or similar viewport
which is used to allow composition or input of a 'complex character',
in other words it is an "input method window." @Since: ATK-1.12.1

### ATK_ROLE_TABLE_ROW = 88
A row in a table.  @Since: ATK-2.1.0

### ATK_ROLE_TREE_ITEM = 89
An object that represents an element of a tree.  @Since: ATK-2.1.0

### ATK_ROLE_DOCUMENT_SPREADSHEET = 90
A document frame which contains a spreadsheet.  @Since: ATK-2.1.0

### ATK_ROLE_DOCUMENT_PRESENTATION = 91
A document frame which contains a presentation or slide content.  @Since: ATK-2.1.0

### ATK_ROLE_DOCUMENT_TEXT = 92
A document frame which contains textual content, such as found in a word processing application.  @Since: ATK-2.1.0

### ATK_ROLE_DOCUMENT_WEB = 93
A document frame which contains HTML or other markup suitable for display in a web browser.  @Since: ATK-2.1.0

### ATK_ROLE_DOCUMENT_EMAIL = 94
A document frame which contains email content to be displayed or composed either in plain text or HTML.  @Since: ATK-2.1.0

### ATK_ROLE_COMMENT = 95
An object found within a document and designed to present a comment, note, or other annotation. In some cases, this object might not be visible until activated.  @Since: ATK-2.1.0

### ATK_ROLE_LIST_BOX = 96
A non-collapsible list of choices the user can select from. @Since: ATK-2.1.0

### ATK_ROLE_GROUPING = 97
A group of related widgets. This group typically has a label. @Since: ATK-2.1.0

### ATK_ROLE_IMAGE_MAP = 98
An image map object. Usually a graphic with multiple hotspots, where each hotspot can be activated resulting in the loading of another document or section of a document. @Since: ATK-2.1.0

### ATK_ROLE_NOTIFICATION = 99
A transitory object designed to present a message to the user, typically at the desktop level rather than inside a particular application.  @Since: ATK-2.1.0

### ATK_ROLE_INFO_BAR = 100
An object designed to present a message to the user within an existing window. @Since: ATK-2.1.0

### ATK_ROLE_LEVEL_BAR = 101
A bar that serves as a level indicator to, for instance, show the strength of a password or the state of a battery.  @Since: ATK-2.7.3

### ATK_ROLE_TITLE_BAR = 102
A bar that serves as the title of a window or a
dialog. @Since: ATK-2.12

### ATK_ROLE_BLOCK_QUOTE = 103
An object which contains a text section
that is quoted from another source. @Since: ATK-2.12

### ATK_ROLE_AUDIO = 104
An object which represents an audio element. @Since: ATK-2.12

### ATK_ROLE_VIDEO = 105
An object which represents a video element. @Since: ATK-2.12

### ATK_ROLE_DEFINITION = 106
A definition of a term or concept. @Since: ATK-2.12

### ATK_ROLE_ARTICLE = 107
A section of a page that consists of a
composition that forms an independent part of a document, page, or
site. Examples: A blog entry, a news story, a forum post. @Since:
ATK-2.12

### ATK_ROLE_LANDMARK = 108
A region of a web page intended as a
navigational landmark. This is designed to allow Assistive
Technologies to provide quick navigation among key regions within a
document. @Since: ATK-2.12

### ATK_ROLE_LOG = 109
A text widget or container holding log content, such
as chat history and error logs. In this role there is a
relationship between the arrival of new items in the log and the
reading order. The log contains a meaningful sequence and new
information is added only to the end of the log, not at arbitrary
points. @Since: ATK-2.12

### ATK_ROLE_MARQUEE = 110
A container where non-essential information
changes frequently. Common usages of marquee include stock tickers
and ad banners. The primary difference between a marquee and a log
is that logs usually have a meaningful order or sequence of
important content changes. @Since: ATK-2.12

### ATK_ROLE_MATH = 111
A text widget or container that holds a mathematical
expression. @Since: ATK-2.12

### ATK_ROLE_RATING = 112
A widget whose purpose is to display a rating,
such as the number of stars associated with a song in a media
player. Objects of this role should also implement
AtkValue. @Since: ATK-2.12

### ATK_ROLE_TIMER = 113
An object containing a numerical counter which
indicates an amount of elapsed time from a start point, or the time
remaining until an end point. @Since: ATK-2.12

### ATK_ROLE_DESCRIPTION_LIST = 114
An object that represents a list of
term-value groups. A term-value group represents a individual
description and consist of one or more names
(ATK_ROLE_DESCRIPTION_TERM) followed by one or more values
(ATK_ROLE_DESCRIPTION_VALUE). For each list, there should not be
more than one group with the same term name. @Since: ATK-2.12

### ATK_ROLE_DESCRIPTION_TERM = 115
An object that represents a term or phrase
with a corresponding definition. @Since: ATK-2.12

### ATK_ROLE_DESCRIPTION_VALUE = 116
An object that represents the
description, definition or value of a term. @Since: ATK-2.12

### ATK_ROLE_STATIC = 117
A generic non-container object whose purpose is to display a
brief amount of information to the user and whose role is known by the
implementor but lacks semantic value for the user. Examples in which
ATK_ROLE_STATIC is appropriate include the message displayed in a message box
and an image used as an alternative means to display text. ATK_ROLE_STATIC
should not be applied to widgets which are traditionally interactive, objects
which display a significant amount of content, or any object which has an
accessible relation pointing to another object. Implementors should expose the
displayed information through the accessible name of the object. If doing so seems
inappropriate, it may indicate that a different role should be used. For
labels which describe another widget, see ATK_ROLE_LABEL. For text views, see
ATK_ROLE_TEXT. For generic containers, see ATK_ROLE_PANEL. For objects whose
role is not known by the implementor, see ATK_ROLE_UNKNOWN. @Since: ATK-2.16.

### ATK_ROLE_MATH_FRACTION = 118
An object that represents a mathematical fraction.

### ATK_ROLE_MATH_ROOT = 119
An object that represents a mathematical expression
displayed with a radical. @Since: ATK-2.16.

### ATK_ROLE_SUBSCRIPT = 120
An object that contains text that is displayed as a
subscript. @Since: ATK-2.16.

### ATK_ROLE_SUPERSCRIPT = 121
An object that contains text that is displayed as a
superscript. @Since: ATK-2.16.

### ATK_ROLE_FOOTNOTE = 122
An object that contains the text of a footnote. @Since: ATK-2.26.

### ATK_ROLE_LAST_DEFINED = 123
not a valid role, used for finding end of the enumeration


## `StateType`

The possible types of states of an object

C - `AtkStateType`

### ATK_STATE_INVALID = 0
Indicates an invalid state - probably an error condition.

### ATK_STATE_ACTIVE = 1
Indicates a window is currently the active window, or an object is the active subelement within a container or table. ATK_STATE_ACTIVE should not be used for objects which have ATK_STATE_FOCUSABLE or ATK_STATE_SELECTABLE: Those objects should use ATK_STATE_FOCUSED and ATK_STATE_SELECTED respectively. ATK_STATE_ACTIVE is a means to indicate that an object which is not focusable and not selectable is the currently-active item within its parent container.

### ATK_STATE_ARMED = 2
Indicates that the object is 'armed', i.e. will be activated by if a pointer button-release event occurs within its bounds.  Buttons often enter this state when a pointer click occurs within their bounds, as a precursor to activation. ATK_STATE_ARMED has been deprecated since ATK-2.16 and should not be used in newly-written code.

### ATK_STATE_BUSY = 3
Indicates the current object is busy, i.e. onscreen representation is in the process of changing, or the object is temporarily unavailable for interaction due to activity already in progress.  This state may be used by implementors of Document to indicate that content loading is underway.  It also may indicate other 'pending' conditions; clients may wish to interrogate this object when the ATK_STATE_BUSY flag is removed.

### ATK_STATE_CHECKED = 4
Indicates this object is currently checked, for instance a checkbox is 'non-empty'.

### ATK_STATE_DEFUNCT = 5
Indicates that this object no longer has a valid backing widget (for instance, if its peer object has been destroyed)

### ATK_STATE_EDITABLE = 6
Indicates that this object can contain text, and that the
user can change the textual contents of this object by editing those contents
directly. For an object which is expected to be editable due to its type, but
which cannot be edited due to the application or platform preventing the user
from doing so, that object's #AtkStateSet should lack ATK_STATE_EDITABLE and
should contain ATK_STATE_READ_ONLY.

### ATK_STATE_ENABLED = 7
Indicates that this object is enabled, i.e. that it currently reflects some application state. Objects that are "greyed out" may lack this state, and may lack the STATE_SENSITIVE if direct user interaction cannot cause them to acquire STATE_ENABLED. See also: ATK_STATE_SENSITIVE

### ATK_STATE_EXPANDABLE = 8
Indicates this object allows progressive disclosure of its children

### ATK_STATE_EXPANDED = 9
Indicates this object its expanded - see ATK_STATE_EXPANDABLE above

### ATK_STATE_FOCUSABLE = 10
Indicates this object can accept keyboard focus, which means all events resulting from typing on the keyboard will normally be passed to it when it has focus

### ATK_STATE_FOCUSED = 11
Indicates this object currently has the keyboard focus

### ATK_STATE_HORIZONTAL = 12
Indicates the orientation of this object is horizontal; used, for instance, by objects of ATK_ROLE_SCROLL_BAR.  For objects where vertical/horizontal orientation is especially meaningful.

### ATK_STATE_ICONIFIED = 13
Indicates this object is minimized and is represented only by an icon

### ATK_STATE_MODAL = 14
Indicates something must be done with this object before the user can interact with an object in a different window

### ATK_STATE_MULTI_LINE = 15
Indicates this (text) object can contain multiple lines of text

### ATK_STATE_MULTISELECTABLE = 16
Indicates this object allows more than one of its children to be selected at the same time, or in the case of text objects, that the object supports non-contiguous text selections.

### ATK_STATE_OPAQUE = 17
Indicates this object paints every pixel within its rectangular region.

### ATK_STATE_PRESSED = 18
Indicates this object is currently pressed.

### ATK_STATE_RESIZABLE = 19
Indicates the size of this object is not fixed

### ATK_STATE_SELECTABLE = 20
Indicates this object is the child of an object that allows its children to be selected and that this child is one of those children that can be selected

### ATK_STATE_SELECTED = 21
Indicates this object is the child of an object that allows its children to be selected and that this child is one of those children that has been selected

### ATK_STATE_SENSITIVE = 22
Indicates this object is sensitive, e.g. to user interaction.
STATE_SENSITIVE usually accompanies STATE_ENABLED for user-actionable controls,
but may be found in the absence of STATE_ENABLED if the current visible state of the
control is "disconnected" from the application state.  In such cases, direct user interaction
can often result in the object gaining STATE_SENSITIVE, for instance if a user makes
an explicit selection using an object whose current state is ambiguous or undefined.
@see STATE_ENABLED, STATE_INDETERMINATE.

### ATK_STATE_SHOWING = 23
Indicates this object, the object's parent, the object's parent's parent, and so on,
are all 'shown' to the end-user, i.e. subject to "exposure" if blocking or obscuring objects do not interpose
between this object and the top of the window stack.

### ATK_STATE_SINGLE_LINE = 24
Indicates this (text) object can contain only a single line of text

### ATK_STATE_STALE = 25
Indicates that the information returned for this object may no longer be
synchronized with the application state.  This is implied if the object has STATE_TRANSIENT,
and can also occur towards the end of the object peer's lifecycle. It can also be used to indicate that
the index associated with this object has changed since the user accessed the object (in lieu of
"index-in-parent-changed" events).

### ATK_STATE_TRANSIENT = 26
Indicates this object is transient, i.e. a snapshot which may not emit events when its
state changes.  Data from objects with ATK_STATE_TRANSIENT should not be cached, since there may be no
notification given when the cached data becomes obsolete.

### ATK_STATE_VERTICAL = 27
Indicates the orientation of this object is vertical

### ATK_STATE_VISIBLE = 28
Indicates this object is visible, e.g. has been explicitly marked for exposure to the user.

### ATK_STATE_MANAGES_DESCENDANTS = 29
Indicates that "active-descendant-changed" event
is sent when children become 'active' (i.e. are selected or navigated to onscreen).
Used to prevent need to enumerate all children in very large containers, like tables.
The presence of STATE_MANAGES_DESCENDANTS is an indication to the client.
that the children should not, and need not, be enumerated by the client.
Objects implementing this state are expected to provide relevant state
notifications to listening clients, for instance notifications of visibility
changes and activation of their contained child objects, without the client
having previously requested references to those children.

### ATK_STATE_INDETERMINATE = 30
Indicates that the value, or some other quantifiable
property, of this AtkObject cannot be fully determined. In the case of a large
data set in which the total number of items in that set is unknown (e.g. 1 of
999+), implementors should expose the currently-known set size (999) along
with this state. In the case of a check box, this state should be used to
indicate that the check box is a tri-state check box which is currently
neither checked nor unchecked.

### ATK_STATE_TRUNCATED = 31
Indicates that an object is truncated, e.g. a text value in a speradsheet cell.

### ATK_STATE_REQUIRED = 32
Indicates that explicit user interaction with an object is required by the user interface, e.g. a required field in a "web-form" interface.

### ATK_STATE_INVALID_ENTRY = 33
Indicates that the object has encountered an error condition due to failure of input validation. For instance, a form control may acquire this state in response to invalid or malformed user input.

### ATK_STATE_SUPPORTS_AUTOCOMPLETION = 34
Indicates that the object in question implements some form of ¨typeahead¨ or
pre-selection behavior whereby entering the first character of one or more sub-elements
causes those elements to scroll into view or become selected.  Subsequent character input
may narrow the selection further as long as one or more sub-elements match the string.
This state is normally only useful and encountered on objects that implement Selection.
In some cases the typeahead behavior may result in full or partial ¨completion¨ of
the data in the input field, in which case these input events may trigger text-changed
events from the AtkText interface.  This state supplants @ATK_ROLE_AUTOCOMPLETE.

### ATK_STATE_SELECTABLE_TEXT = 35
Indicates that the object in question supports text selection. It should only be exposed on objects which implement the Text interface, in order to distinguish this state from @ATK_STATE_SELECTABLE, which infers that the object in question is a selectable child of an object which implements Selection. While similar, text selection and subelement selection are distinct operations.

### ATK_STATE_DEFAULT = 36
Indicates that the object is the "default" active component, i.e. the object which is activated by an end-user press of the "Enter" or "Return" key.  Typically a "close" or "submit" button.

### ATK_STATE_ANIMATED = 37
Indicates that the object changes its appearance dynamically as an inherent part of its presentation.  This state may come and go if an object is only temporarily animated on the way to a 'final' onscreen presentation.
@note some applications, notably content viewers, may not be able to detect
all kinds of animated content.  Therefore the absence of this state should not
be taken as definitive evidence that the object's visual representation is
static; this state is advisory.

### ATK_STATE_VISITED = 38
Indicates that the object (typically a hyperlink) has already been 'activated', and/or its backing data has already been downloaded, rendered, or otherwise "visited".

### ATK_STATE_CHECKABLE = 39
Indicates this object has the potential to be
 checked, such as a checkbox or toggle-able table cell. @Since:
 ATK-2.12

### ATK_STATE_HAS_POPUP = 40
Indicates that the object has a popup context
menu or sub-level menu which may or may not be showing. This means
that activation renders conditional content.  Note that ordinary
tooltips are not considered popups in this context. @Since: ATK-2.12

### ATK_STATE_HAS_TOOLTIP = 41
Indicates this object has a tooltip. @Since: ATK-2.16

### ATK_STATE_READ_ONLY = 42
Indicates that a widget which is ENABLED and SENSITIVE
has a value which can be read, but not modified, by the user. Note that this
state should only be applied to widget types whose value is normally directly
user modifiable, such as check boxes, radio buttons, spin buttons, text input
fields, and combo boxes, as a means to convey that the expected interaction
with that widget is not possible. When the expected interaction with a
widget does not include modification by the user, as is the case with
labels and containers, ATK_STATE_READ_ONLY should not be applied. See also
ATK_STATE_EDITABLE. @Since: ATK-2-16

### ATK_STATE_LAST_DEFINED = 43
Not a valid state, used for finding end of enumeration


## `TextAttribute`

Describes the text attributes supported

C - `AtkTextAttribute`

### ATK_TEXT_ATTR_INVALID = 0
Invalid attribute, like bad spelling or grammar.

### ATK_TEXT_ATTR_LEFT_MARGIN = 1
The pixel width of the left margin

### ATK_TEXT_ATTR_RIGHT_MARGIN = 2
The pixel width of the right margin

### ATK_TEXT_ATTR_INDENT = 3
The number of pixels that the text is indented

### ATK_TEXT_ATTR_INVISIBLE = 4
Either "true" or "false" indicating whether text is visible or not

### ATK_TEXT_ATTR_EDITABLE = 5
Either "true" or "false" indicating whether text is editable or not

### ATK_TEXT_ATTR_PIXELS_ABOVE_LINES = 6
Pixels of blank space to leave above each newline-terminated line.

### ATK_TEXT_ATTR_PIXELS_BELOW_LINES = 7
Pixels of blank space to leave below each newline-terminated line.

### ATK_TEXT_ATTR_PIXELS_INSIDE_WRAP = 8
Pixels of blank space to leave between wrapped lines inside the same newline-terminated line (paragraph).

### ATK_TEXT_ATTR_BG_FULL_HEIGHT = 9
"true" or "false" whether to make the background color for each character the height of the highest font used on the current line, or the height of the font used for the current character.

### ATK_TEXT_ATTR_RISE = 10
Number of pixels that the characters are risen above the baseline

### ATK_TEXT_ATTR_UNDERLINE = 11
"none", "single", "double", "low", or "error"

### ATK_TEXT_ATTR_STRIKETHROUGH = 12
"true" or "false" whether the text is strikethrough

### ATK_TEXT_ATTR_SIZE = 13
The size of the characters in points. eg: 10

### ATK_TEXT_ATTR_SCALE = 14
The scale of the characters. The value is a string representation of a double

### ATK_TEXT_ATTR_WEIGHT = 15
The weight of the characters.

### ATK_TEXT_ATTR_LANGUAGE = 16
The language used

### ATK_TEXT_ATTR_FAMILY_NAME = 17
The font family name

### ATK_TEXT_ATTR_BG_COLOR = 18
The background color. The value is an RGB value of the format "%u,%u,%u"

### ATK_TEXT_ATTR_FG_COLOR = 19
The foreground color. The value is an RGB value of the format "%u,%u,%u"

### ATK_TEXT_ATTR_BG_STIPPLE = 20
"true" if a #GdkBitmap is set for stippling the background color.

### ATK_TEXT_ATTR_FG_STIPPLE = 21
"true" if a #GdkBitmap is set for stippling the foreground color.

### ATK_TEXT_ATTR_WRAP_MODE = 22
The wrap mode of the text, if any. Values are "none", "char", "word", or "word_char".

### ATK_TEXT_ATTR_DIRECTION = 23
The direction of the text, if set. Values are "none", "ltr" or "rtl"

### ATK_TEXT_ATTR_JUSTIFICATION = 24
The justification of the text, if set. Values are "left", "right", "center" or "fill"

### ATK_TEXT_ATTR_STRETCH = 25
The stretch of the text, if set. Values are "ultra_condensed", "extra_condensed", "condensed", "semi_condensed", "normal", "semi_expanded", "expanded", "extra_expanded" or "ultra_expanded"

### ATK_TEXT_ATTR_VARIANT = 26
The capitalization variant of the text, if set. Values are "normal" or "small_caps"

### ATK_TEXT_ATTR_STYLE = 27
The slant style of the text, if set. Values are "normal", "oblique" or "italic"

### ATK_TEXT_ATTR_LAST_DEFINED = 28
not a valid text attribute, used for finding end of enumeration


## `TextBoundary`

Text boundary types used for specifying boundaries for regions of text.
This enumeration is deprecated since 2.9.4 and should not be used. Use
AtkTextGranularity with #atk_text_get_string_at_offset instead.

C - `AtkTextBoundary`

### ATK_TEXT_BOUNDARY_CHAR = 0
Boundary is the boundary between characters
(including non-printing characters)

### ATK_TEXT_BOUNDARY_WORD_START = 1
Boundary is the start (i.e. first character) of a word.

### ATK_TEXT_BOUNDARY_WORD_END = 2
Boundary is the end (i.e. last
character) of a word.

### ATK_TEXT_BOUNDARY_SENTENCE_START = 3
Boundary is the first character in a sentence.

### ATK_TEXT_BOUNDARY_SENTENCE_END = 4
Boundary is the last (terminal)
character in a sentence; in languages which use "sentence stop"
punctuation such as English, the boundary is thus the '.', '?', or
similar terminal punctuation character.

### ATK_TEXT_BOUNDARY_LINE_START = 5
Boundary is the initial character of the content or a
character immediately following a newline, linefeed, or return character.

### ATK_TEXT_BOUNDARY_LINE_END = 6
Boundary is the linefeed, or return
character.


## `TextClipType`

Describes the type of clipping required.

C - `AtkTextClipType`

### ATK_TEXT_CLIP_NONE = 0
No clipping to be done

### ATK_TEXT_CLIP_MIN = 1
Text clipped by min coordinate is omitted

### ATK_TEXT_CLIP_MAX = 2
Text clipped by max coordinate is omitted

### ATK_TEXT_CLIP_BOTH = 3
Only text fully within mix/max bound is retained


## `TextGranularity`

Text granularity types used for specifying the granularity of the region of
text we are interested in.

C - `AtkTextGranularity`

### ATK_TEXT_GRANULARITY_CHAR = 0
Granularity is defined by the boundaries between characters
(including non-printing characters)

### ATK_TEXT_GRANULARITY_WORD = 1
Granularity is defined by the boundaries of a word,
starting at the beginning of the current word and finishing at the beginning of
the following one, if present.

### ATK_TEXT_GRANULARITY_SENTENCE = 2
Granularity is defined by the boundaries of a sentence,
starting at the beginning of the current sentence and finishing at the beginning of
the following one, if present.

### ATK_TEXT_GRANULARITY_LINE = 3
Granularity is defined by the boundaries of a line,
starting at the beginning of the current line and finishing at the beginning of
the following one, if present.

### ATK_TEXT_GRANULARITY_PARAGRAPH = 4
Granularity is defined by the boundaries of a paragraph,
starting at the beginning of the current paragraph and finishing at the beginning of
the following one, if present.


## `ValueType`

Default types for a given value. Those are defined in order to
easily get localized strings to describe a given value or a given
subrange, using atk_value_type_get_localized_name().

C - `AtkValueType`

### ATK_VALUE_VERY_WEAK = 0


### ATK_VALUE_WEAK = 1


### ATK_VALUE_ACCEPTABLE = 2


### ATK_VALUE_STRONG = 3


### ATK_VALUE_VERY_STRONG = 4


### ATK_VALUE_VERY_LOW = 5


### ATK_VALUE_LOW = 6


### ATK_VALUE_MEDIUM = 7


### ATK_VALUE_HIGH = 8


### ATK_VALUE_VERY_HIGH = 9


### ATK_VALUE_VERY_BAD = 10


### ATK_VALUE_BAD = 11


### ATK_VALUE_GOOD = 12


### ATK_VALUE_VERY_GOOD = 13


### ATK_VALUE_BEST = 14


### ATK_VALUE_LAST_DEFINED = 15


