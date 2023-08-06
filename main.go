package main

import "github.com/DavidEsdrs/ads/graphs"

func main() {
	graph := graphs.NewGraph[string]()

	graph.CreateEdge("A", "B", "C", "D")
	graph.CreateEdge("B", "E", "F")
	graph.CreateEdge("C", "G", "H", "I")
	graph.CreateEdge("D", "J", "K", "L")
	graph.CreateEdge("E", "M", "N")
	graph.CreateEdge("F", "O", "P", "Q")
	graph.CreateEdge("G", "R", "S")
	graph.CreateEdge("H", "T", "U")
	graph.CreateEdge("I", "V", "W")
	graph.CreateEdge("J", "X", "Y", "Z")
	graph.CreateEdge("K", "AA", "AB", "AC")
	graph.CreateEdge("L", "AD", "AE", "AF")
	graph.CreateEdge("M", "AG", "AH", "AI")
	graph.CreateEdge("N", "AJ", "AK", "AL")
	graph.CreateEdge("O", "AM", "AN", "AO")
	graph.CreateEdge("P", "AP", "AQ", "AR")
	graph.CreateEdge("Q", "AS", "AT", "AU")
	graph.CreateEdge("R", "AV", "AW", "AX")
	graph.CreateEdge("S", "AY", "AZ", "BA")
	graph.CreateEdge("T", "BB", "BC", "BD")
	graph.CreateEdge("U", "BE", "BF", "BG")
	graph.CreateEdge("V", "BH", "BI", "BJ")
	graph.CreateEdge("W", "BK", "BL", "BM")
	graph.CreateEdge("X", "BN", "BO", "BP")
	graph.CreateEdge("Y", "BQ", "BR", "BS")
	graph.CreateEdge("Z", "BT", "BU", "BV")
	graph.CreateEdge("AA", "BW", "BX", "BY")
	graph.CreateEdge("AB", "BZ", "CA", "CB")
	graph.CreateEdge("AC", "CC", "CD", "CE")
	graph.CreateEdge("AD", "CF", "CG", "CH")
	graph.CreateEdge("AE", "CI", "CJ", "CK")
	graph.CreateEdge("AF", "CL", "CM", "CN")
	graph.CreateEdge("AG", "CO", "CP", "CQ")
	graph.CreateEdge("AH", "CR", "CS", "CT")
	graph.CreateEdge("AI", "CU", "CV", "CW")
	graph.CreateEdge("AJ", "CX", "CY", "CZ")
	graph.CreateEdge("AK", "DA", "DB", "DC")
	graph.CreateEdge("AL", "DD", "DE", "DF")
	graph.CreateEdge("AM", "DG", "DH", "DI")
	graph.CreateEdge("AN", "DJ", "DK", "DL")
	graph.CreateEdge("AO", "DM", "DN", "DO")
	graph.CreateEdge("AP", "DP", "DQ", "DR")
	graph.CreateEdge("AQ", "DS", "DT", "DU")
	graph.CreateEdge("AR", "DV", "DW", "DX")
	graph.CreateEdge("AS", "DY", "DZ", "EA")
	graph.CreateEdge("AT", "EB", "EC", "ED")
	graph.CreateEdge("AU", "EE", "EF", "EG")
	graph.CreateEdge("AV", "EH", "EI", "EJ")
	graph.CreateEdge("AW", "EK", "EL", "EM")
	graph.CreateEdge("AX", "EN", "EO", "EP")
	graph.CreateEdge("AY", "EQ", "ER", "ES")
	graph.CreateEdge("AZ", "ET", "EU", "EV")
	graph.CreateEdge("BA", "EW", "EX", "EY")
	graph.CreateEdge("BB", "EZ", "FA", "FB")
	graph.CreateEdge("BC", "FC", "FD", "FE")
	graph.CreateEdge("BD", "FF", "FG", "FH")
	graph.CreateEdge("BE", "FI", "FJ", "FK")
	graph.CreateEdge("BF", "FL", "FM", "FN")
	graph.CreateEdge("BG", "FO", "FP", "FQ")
	graph.CreateEdge("BH", "FR", "FS", "FT")
	graph.CreateEdge("BI", "FU", "FV", "FW")

	graphs.DepthFirst(&graph, "CO")
	println(graph.CountComponents())
	println(graph.HasPath("A", "CO"))
}
