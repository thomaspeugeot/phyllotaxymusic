package models

import (
	"fmt"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/musicxml"
)

func (*Parameter) generateHeaders(scorePartwise *m.Score_partwise) {
	scorePartwise.XMLName.Local = "score-partwise"
	scorePartwise.A_score_partwise.Version = "4.0"

	// timestamp := time.Now().Format("20060102_150405")
	// scorePartwise.Work =
	// 	&m.Work{

	// 		Name:        "Phyllotaxy music",
	// 		Work_number: "",
	// 		Work_title:  "Phyllotaxy music " + timestamp,
	// 		Opus: &m.Opus{
	// 			Name: "",
	// 		},
	// 	}
	{
		var identification m.Identification
		scorePartwise.Identification = &identification

		var creator m.Typed_text
		identification.Creator = append(identification.Creator, &creator)
		creator.EnclosedText = "phyllotaxy music"
		creator.Type = "composer"

		var encoding m.Encoding
		identification.Encoding = &encoding

		encoding.Software = "MuseScore 4.3.2"

		{
			var supports m.Supports
			encoding.Supports = append(encoding.Supports, &supports)

			supports.Type = m.Enum_Yes_no_Yes
			supports.Element = "accidental"
		}
		{
			var supports m.Supports
			encoding.Supports = append(encoding.Supports, &supports)

			supports.Type = m.Enum_Yes_no_Yes
			supports.Element = "beam"
		}
		{
			var supports m.Supports
			encoding.Supports = append(encoding.Supports, &supports)

			supports.Type = m.Enum_Yes_no_Yes
			supports.Element = "print"
			supports.Attribute = "new-page"
			supports.Value = string(m.Enum_Yes_no_Yes)

		}
		{
			var supports m.Supports
			encoding.Supports = append(encoding.Supports, &supports)

			supports.Type = m.Enum_Yes_no_Yes
			supports.Element = "print"
			supports.Attribute = "new-system"
			supports.Value = string(m.Enum_Yes_no_Yes)
		}
		{
			var supports m.Supports
			encoding.Supports = append(encoding.Supports, &supports)

			supports.Type = m.Enum_Yes_no_Yes
			supports.Element = "stem"
		}
	}

	{
		var defaults m.Defaults
		scorePartwise.Defaults = &defaults

		{
			var scaling m.Scaling
			defaults.Scaling = &scaling

			// <millimeters>: Specifies the size of one staff space in millimeters.
			// In this example, one staff space equals approximately 6.99912 mm.

			// <tenths>: Defines the number of tenths that correspond to one staff space.
			// Here, 40 tenths equals one staff space.
			// This value is commonly used for positioning notes and symbols in the score.

			// A staff space refers to the vertical distance between two adjacent
			// lines of a musical staff (or stave). In musical notation,
			// the staff is a set of five horizontal
			// lines on which notes and other musical symbols are written.
			scaling.Millimeters = fmt.Sprintf("%f", 6.99912)
			scaling.Tenths = fmt.Sprintf("%d", 40)
		}
		{
			var page_layout m.Page_layout
			defaults.Page_layout = &page_layout

			// A4 paper approx.
			page_layout.Page_height = fmt.Sprintf("%f", 1596.77)
			page_layout.Page_width = fmt.Sprintf("%f", 1233.87)

			{
				var page_margins m.Page_margins
				page_layout.Page_margins = &page_margins

				page_margins.Left_margin = fmt.Sprintf("%f", 85.725)
				page_margins.Right_margin = fmt.Sprintf("%f", 85.725)
				page_margins.Top_margin = fmt.Sprintf("%f", 85.725)
				page_margins.Bottom_margin = fmt.Sprintf("%f", 85.725)
			}

		}
		// {
		// 	var appearance m.Appearance
		// 	defaults.Appearance = &appearance

		// 	appearance.Line_width = append(appearance.Line_width,
		// 		&m.Line_width{Type: "heavy barline", EnclosedText: "5.5"},
		// 		&m.Line_width{Type: "beam", EnclosedText: "5"},
		// 		&m.Line_width{Type: "bracket", EnclosedText: "4.5"},
		// 		&m.Line_width{Type: "dashes", EnclosedText: "1"},
		// 		&m.Line_width{Type: "enclosure", EnclosedText: "1"},
		// 		&m.Line_width{Type: "ending", EnclosedText: "1.1"},
		// 		&m.Line_width{Type: "extend", EnclosedText: "1"},
		// 		&m.Line_width{Type: "leger", EnclosedText: "1.6"},
		// 		&m.Line_width{Type: "pedal", EnclosedText: "1.1"},
		// 		&m.Line_width{Type: "octave shift", EnclosedText: "1.1"},
		// 		&m.Line_width{Type: "slur middle", EnclosedText: "2.1"},
		// 		&m.Line_width{Type: "slur tip", EnclosedText: "0.5"},
		// 		&m.Line_width{Type: "staff", EnclosedText: "1.1"},
		// 		&m.Line_width{Type: "stem", EnclosedText: "1"},
		// 		&m.Line_width{Type: "tie middle", EnclosedText: "2.1"},
		// 		&m.Line_width{Type: "tie tip", EnclosedText: "0.5"},
		// 		&m.Line_width{Type: "tuplet bracket", EnclosedText: "1"},
		// 		&m.Line_width{Type: "wedge", EnclosedText: "1.2"},
		// 	)

		// 	appearance.Note_size = append(appearance.Note_size,
		// 		&m.Note_size{Type: "cue", EnclosedText: "70"},
		// 		&m.Note_size{Type: "grace", EnclosedText: "70"},
		// 		&m.Note_size{Type: "grace-cue", EnclosedText: "49"},
		// 	)
		// }
		defaults.Music_font = &m.Empty_font{
			AttributeGroup_font: m.AttributeGroup_font{
				Font_family: "Times New Roman",
			},
		}

		defaults.Word_font = &m.Empty_font{
			AttributeGroup_font: m.AttributeGroup_font{
				Font_family: "Times New Roman",
			},
		}

	}

	{
		var credit m.Credit
		credit.Page = 1

		var formatedText m.Formatted_text_id
		formatedText.EnclosedText = "Phyllotaxy music"
		formatedText.Default_x = "595.44"
		formatedText.Default_y = "1626.67"
		formatedText.Justify = m.Enum_Left_center_right_Center
		formatedText.Font_size = "24"
		formatedText.Valign = "top"

		credit.Credit_words = append(credit.Credit_words, &formatedText)
		scorePartwise.Credit = append(scorePartwise.Credit, &credit)
	}
	{
		var part_list m.Part_list
		scorePartwise.Part_list = &part_list

		{
			var score_part m.Score_part
			part_list.Score_part = &score_part

			score_part.Id = "P1"

			score_part.Part_name = &m.Part_name{EnclosedText: "Piano"}
			score_part.Part_abbreviation = &m.Part_name{EnclosedText: "Pno."}

			{
				var score_instrument m.Score_instrument
				score_part.Score_instrument = append(score_part.Score_instrument, &score_instrument)
				score_instrument.Instrument_name = "Piano"
				score_instrument.Instrument_sound = "keyboard.piano"
			}
			{
				score_part.Midi_device = append(score_part.Midi_device,
					&m.Midi_device{
						Port: 1,
						Id:   "P1-I1",
					})
			}
			{
				var midi_instrument m.Midi_instrument
				score_part.Midi_instrument = append(score_part.Midi_instrument, &midi_instrument)

				midi_instrument.Midi_channel = 1
				midi_instrument.Midi_program = 1
				midi_instrument.Volume = fmt.Sprintf("%f", 78.7402)
				midi_instrument.Pan = "0"
			}
		}
	}
}
