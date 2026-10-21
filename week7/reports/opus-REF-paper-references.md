# Reference validation: Gaver, "What Should We Expect From Research Through Design?", CHI 2012, pp. 937 to 946

Checked 2026-10-19 against the PDF (reference list on printed p. 946, PDF page 10) and `assets/paper-text.txt`. Sources: Crossref API (`https://api.crossref.org/works/<DOI>` or `?query.bibliographic=`), publisher pages, Open Library and Internet Archive records, and the authors' or journals' own PDFs. Where a printed entry contains an en-dash in a page range, it is shown here with a hyphen.

Verdicts: **OK** (data matches the source; small omissions such as missing pages are noted but not counted as errors), **wrong** (a factual error in the entry), **unverifiable** (no source found).

## 1. The 36 reference-list entries

| # | Entry as printed (p. 946) | What the source says | URL used | Verdict |
|---|---|---|---|---|
| 1 | Alexander C, Ishikawa S (1977) A pattern language: towns, buildings, construction. Oxford University Press. | Alexander, C., Ishikawa, S., Silverstein, M. (with Jacobson, Fiksdahl-King, Angel), 1977, Oxford University Press, New York. | https://openlibrary.org/works/OL3923600W | wrong (minor): third principal author Murray Silverstein omitted; title, year, publisher correct |
| 2 | Bardzell J, Bardzell S, DiSalvo C, Gaver B and Sengers P (2012). Panel: The humanities and/in HCI. CHI'2012. | "The humanities and/in HCI", Bardzell, Bardzell, DiSalvo, Gaver, Sengers, CHI '12 Extended Abstracts, pp. 1135-1138. | https://api.crossref.org/works/10.1145/2212776.2212405 | OK (pages omitted) |
| 3 | Collins R (1994). Why the Social Sciences Won't Become High-Consensus, Rapid-Discovery Science. Sociological Forum 9(2) 155-177. | Collins, R., Sociological Forum 9(2), 155-177, June 1994. | https://api.crossref.org/works/10.1007/bf01476360 | OK as an entry (but see section 2: three in-text uses of [3] mean another work) |
| 4 | Crampton-Smith, G. and Tabor, P. 1996. The role of the artist-designer. In Bringing Design To Software, T. Winograd, Ed. ACM, New York, NY, 37-61. | "The role of the artist-designer", Crampton Smith and Tabor, in Bringing Design to Software, April 1996, pp. 37-61. | https://api.crossref.org/works/10.1145/229868.230031 | OK |
| 5 | Curd M and Cover J (1998). Philosophy of Science: The Central Issues. London: WW. Norton & Co. | Curd, M. and Cover, J.A. (editors), Philosophy of Science: The Central Issues, 1st ed., New York: W.W. Norton, 1998. | https://openlibrary.org/isbn/0393971759.json ; https://archive.org/metadata/philosophyofscie0000unse_k0g2 | wrong (minor): place of publication is New York, not London; "WW." should be "W.W."; Curd and Cover are editors (anthology) |
| 6 | Djajadiningrat T, Matthews B, Stienstra M. (2007). Easy doesn't do it: skill and expression in tangible aesthetics. Pers Ubiquit Comput 11:657-676. | Personal and Ubiquitous Computing 11(8), 657-676, 2007. | https://api.crossref.org/works?query.bibliographic=Djajadiningrat+Easy+doesn't+do+it (DOI 10.1007/s00779-006-0137-9) | OK (issue 8 omitted) |
| 7 | Fallmann D (2003). Design-oriented human-computer interaction. CHI'2003, 225-232. | Author is Daniel **Fallman**; CHI 2003, pp. 225-232. | https://api.crossref.org/works?query.bibliographic=Fallman+Design-oriented+human-computer+interaction (DOI 10.1145/642611.642652) | wrong: author surname misspelt "Fallmann", correct "Fallman" |
| 8 | Feynman, R. (1989). The character of physical law. London: MIT Press. | First published 1965 (BBC); MIT Press paperback 1967, many later printings, MIT Press Cambridge, Mass. and London. | https://openlibrary.org/works/OL271750W | OK (1989 is a later printing, not the first edition; a 1989 printing itself was not confirmed) |
| 9 | Forlizzi J, DiSalvo C, Bardzell J, Koskinen I, and Wensveen S. (2011). Quality control: A panel on the critique and criticism of design research. CHI 2011 Panel, 823-826. | "Quality control: a panel on the critique and criticism of design research", CHI '11 Extended Abstracts, pp. 823-826. | https://api.crossref.org/works/10.1145/1979742.1979497 | OK |
| 10 | Forlizzi, J. (2008). The Product Ecology: understanding social product use and supporting design culture. International Journal of Design 2 (1), 11-20. | IJD Vol. 2 No. 1 (2008); the article PDF's own citation line gives 2(1), 11-20. | https://www.ijdesign.org/index.php/IJDesign/issue/view/8/showToc ; https://www.ijdesign.org/index.php/IJDesign/article/download/220/136 | OK |
| 11 | Frayling, C. Research in Art and Design. Royal College of Art Research Papers 1, 1 (1993),1-5. | RCA Research Papers Vol 1, No 1, 1993/4, pp. 1-5. | https://researchonline.rca.ac.uk/384/ | OK |
| 12 | Gaver W. (2011) Making spaces: how design workbooks work. Proc. CHI '11, 1551-1560. | CHI 2011, pp. 1551-1560. | https://api.crossref.org/works/10.1145/1978942.1979169 | OK |
| 13 | Gaver, W., (2009). Designing for Homo Ludens, Still. In (Re)searching the Digital Bauhaus. Binder, T., Löwgren, J., and Malmborg, L. (eds.). London: Springer, pp. 163-178. | Chapter in (Re)Searching the Digital Bauhaus, eds. Binder, Löwgren, Malmborg, Springer London, HCI Series, 2009, pp. 163-178 (DOI 10.1007/978-1-84800-350-7_9). | https://api.crossref.org/works/10.1007/978-1-84800-350-7_9 ; https://pure.itu.dk/en/publications/bill-gaver-designing-for-homo-ludens-still/ | OK |
| 14 | Gaver, W., Beaver, J., and Benford, S. (2003). Ambiguity as a resource for design. Proc. CHI'03, Ft. Lauderdale. New York: ACM Press. | CHI 2003, pp. 233-240. | https://api.crossref.org/works/10.1145/642611.642653 (via query) | OK (pages 233-240 omitted) |
| 15 | Gibson J. (1979). The ecological approach to visual perception. Boston: Haughton Mifflin, 1979. | Gibson, J.J., 1979, Boston: **Houghton** Mifflin. | https://openlibrary.org/works/OL5589051W | wrong (minor): publisher misspelt "Haughton", correct "Houghton"; year printed twice |
| 16 | Isbister K, and Höök K. (2007). Supple interfaces: designing and evaluating for richer human connections and experiences. Proc CHI'07 extended abstracts, 2853-2856. | CHI '07 Extended Abstracts, pp. 2853-2856. | https://api.crossref.org/works/10.1145/1240866.1241094 | OK |
| 17 | Kuhn, T (1970). The Structure of Scientific Revolutions. London: University of Chicago Press. | 2nd ed. 1970, University of Chicago Press (Chicago and London). | https://openlibrary.org/works/OL3259254W | OK |
| 18 | Kuhn, T (1977). Objectivity, value judgment, and theory choice. Reprinted in [3]. | Essay (The Essential Tension, 1977) reprinted in Curd and Cover 1998, Part 2. Reference 3 is Collins 1994, a journal article. | https://zolaist.org/wiki/index.php/Philosophy_of_Science:_The_Central_Issues | wrong: "Reprinted in [3]" should be "Reprinted in [5]" |
| 19 | Lakatos I (1977). Science and pseudoscience. Reprinted in [3]. | Reprinted in Curd and Cover 1998, starting p. 20. | https://afterall.net/?p=492056 ; https://zolaist.org/wiki/index.php/Philosophy_of_Science:_The_Central_Issues | wrong: "Reprinted in [3]" should be "Reprinted in [5]" |
| 20 | Latour, B. (1987). Science in Action. Cambridge, Mass: Harvard University Press. | Harvard University Press, Cambridge, Mass., 1987. | https://openlibrary.org/works/OL1998743W | OK |
| 21 | Law, J. and Urry, J. (2004). Enacting the social. Economy and Society 33(3) pp 390-410. | Economy and Society 33(3), 390-410, Aug 2004. | https://api.crossref.org/works?query.bibliographic=Law+Urry+Enacting+the+social (DOI 10.1080/0308514042000225716) | OK |
| 22 | Osborne, T. and Rose, N. (1999). Do the social sciences create phenomena?: the example of public opinion research. British Journal of Sociology 50(3) pp 367-396. | British Journal of Sociology 50(3), 367-396, Sept 1999. | https://api.crossref.org/works?query.bibliographic=Osborne+Rose+Do+the+social+sciences+create+phenomena (DOI 10.1111/j.1468-4446.1999.00367.x) | OK |
| 23 | Overbeeke K and Wesveen, S. (2003). From perception to experience, from affordances to irresistibles. Proc. DPPI'03, 92-97. | Overbeeke and **Wensveen**, "From perception to experience, from affordances to irresistibles", DPPI '03, pp. 92-97. | https://api.crossref.org/works?query.bibliographic=Overbeeke+Wensveen+irresistibles (DOI 10.1145/782896.782919) | wrong: second author misspelt "Wesveen", correct "Wensveen" (the body text on p. 938 spells it correctly). Title spelling "irresistibles" is correct |
| 24 | Popper K (1963). Science: conjectures and refutations. Reprinted in [3]. | Reprinted in Curd and Cover 1998, starting p. 3. | https://afterall.net/?p=492056 | wrong: "Reprinted in [3]" should be "Reprinted in [5]" |
| 25 | Pullin, Graham (2010) Curating and creating design collections. Design and Culture 2 (1), 309-328. | "Curating and creating design collections, from Social Mobiles to the Museum of Lost Interactions and Six Speaking Chairs", Design and Culture **2(3)**, 309-328, Nov 2010. | https://api.crossref.org/works?query.bibliographic=Pullin+Curating+and+creating+design+collections (DOI 10.2752/175470710x12789399279877) | wrong: issue is 3, not 1; title is truncated (subtitle omitted) |
| 26 | Rittel H, and Webber M. (1974). Dilemmas in general theory of planning. Design Research and Methods, 8(1), 31-39. | Title is "Dilemmas in **a** General Theory of Planning". Original: Policy Sciences 4(2), 155-169, 1973. The 1974 Design Research and Methods 8(1) reprint could not be confirmed. | https://api.crossref.org/works/10.1007/bf01405730 (via query) | wrong: article "a" missing from the title; the cited 1974 venue is unverifiable (original is Policy Sciences 4(2), 155-169, 1973) |
| 27 | Rogers, Y. (2004). New theoretical approaches for human-computer interaction. In B. Cronin (Vol. Ed.), Annual review of information, science and technology: Vol. 38 (pp. 87-143). Medford, NJ: Information Today. | Annual Review of Information Science and Technology 38(1), 87-143, 2004. | https://api.crossref.org/works?query.bibliographic=Rogers+New+theoretical+approaches (DOI 10.1002/aris.1440380103) | OK (stray comma in the series title) |
| 28 | Schoen, D (1983). The reflective practitioner, how professionals think. New York, Basic Books. | Schön, D.A., The Reflective Practitioner: How Professionals Think in Action, Basic Books, New York, 1983. | https://openlibrary.org/works/OL3466056W | OK ("Schoen" is a transliteration; subtitle shortened). Never cited in the text |
| 29 | Sengers P (2006). Must design become 'scientific'? DIS'06 Workshop on Exploring Design as a Research Activity. | The DIS 2006 workshop "Exploring Design as a Research Activity" exists (Sas and Dix list a paper there), but no record of this Sengers position paper was found. | https://www.research.lancs.ac.uk/portal/en/publications/exploring-the-design-space(215b66f8-aeba-4d25-900e-fa8b90df1cac).html | unverifiable |
| 30 | Sengers P, Boehner K, David S, and Kaye J. (2005). Reflective design. Proc. AARHUS'05, 49-58. 24 | Sengers, Boehner, David, Kaye, "Reflective design", Aarhus 2005 (4th decennial conference on Critical Computing), pp. 49-58. | https://api.crossref.org/works?query.bibliographic=Sengers+Boehner+David+Kaye+Reflective+design (DOI 10.1145/1094562.1094569) | OK (a stray "24" is printed after the pages, a typesetting slip) |
| 31 | Stolterman, E. (2008). The nature of design practice and implications for interaction design research. International Journal of Design 2(1), 55-65. | IJD Vol. 2 No. 1 (2008); article PDF starts on p. 55, 11 pages (55 to 65). The quotes used at p. 59 and p. 61 are on those pages. | https://www.ijdesign.org/index.php/IJDesign/article/view/240 ; https://www.ijdesign.org/index.php/IJDesign/article/download/240/139 | OK |
| 32 | Wiltse, H. and Stolterman, E. (2010). Architectures of interaction: An architectural perspective on digital experience. Proc. NordiCHI 2010, 821-824. | NordiCHI 2010, pp. 821-824. | https://api.crossref.org/works/10.1145/1868914.1869038 | OK |
| 33 | Zimmerman J and Forlizzi J. (2008) The Role of Design Artifacts in Design Theory Construction. Human-Computer Interaction Institute. Paper 37. | CMU repository record dated 2008 (DOI 10.1184/R1/6470537); published version: Artifact 2(1), 41-45, April 2008. | https://api.figshare.com/v2/articles/6470537 (KiltHub, CMU) ; https://api.crossref.org/works?query.bibliographic=Zimmerman+Forlizzi+Role+of+Design+Artifacts (DOI 10.1080/17493460802276893) | OK (the series number "Paper 37" could not be confirmed: repository.cmu.edu refused the connection; the journal version is Artifact 2(1), 41-45) |
| 34 | Zimmerman J, Forlizzi J, and Evenson, S. (2007). Research through design as a method for interaction design research in HCI. Proc. CHI'07 493-502. | CHI 2007, pp. 493-502. | https://api.crossref.org/works?query.bibliographic=Zimmerman+Forlizzi+Evenson (DOI 10.1145/1240624.1240704) | OK |
| 35 | Zimmerman J, Stolterman E, and Forlizzi J (2010). An analysis and critique of research through design: toward a formalization of a research approach. DIS 2010, 310-319. | Title "An analysis and critique of Research through Design", subtitle "**towards** a formalization of a research approach", DIS 2010, pp. 310-319. | https://api.crossref.org/works/10.1145/1858171.1858228 | wrong (minor): subtitle reads "towards", not "toward" |
| 36 | Zimmerman, J. (2009). Designing for the Self: Making products that help people become the person they desire to be. Proc. CHI 2009, 395-404. | CHI 2009, pp. 395-404. | https://api.crossref.org/works/10.1145/1518701.1518765 | OK |

Tally: 36 checked; 24 OK; 11 wrong (1, 5, 7, 15, 18, 19, 23, 24, 25, 26, 35); 1 unverifiable (29).

## 2. In-text citation mapping (pp. 937 to 946)

Every bracketed number in the body was located in `paper-text.txt` and checked against the sentence.

| Page | Citation and context | Points at | Verdict |
|---|---|---|---|
| 937 | "nature and standards of research through design [9, 31, 33, 35]" | panel 2011, Stolterman 2008, Z and F 2008, Z, S and F 2010 | OK |
| 937 | "'integrating design research methods, approaches, and outcomes in HCI' [9]"; "([9] p824)" | CHI 2011 panel, pp. 823-826 | OK (p. 824 inside range) |
| 937 | "'like any other research approach' ([35] p317)"; "([35] p316)" | DIS 2010, pp. 310-319 | OK (pages inside range) |
| 937 | "'a level of extensibility and verifiability' ... 'theory designers can apply in research in practice' [33]" | Z and F 2008 | Points correctly. Both phrases are in the source (CMU PDF), but the source reads "apply in research **and** practice"; the paper's "research in practice" is a small misquotation |
| 937 | "research through design [4, 11, 34]" | Crampton Smith and Tabor, Frayling, Z, F and E 2007 | OK |
| 938 | "the humanities [2]"; "a design discipline [7]" | humanities panel 2012; Fallman 2003 | OK |
| 938 | "affordances, context and situatedness [27] (cited in [31])" | Rogers 2004 cited in Stolterman 2008 | OK (Stolterman p. 56 cites Rogers 2004 for "affordance, context, situatedness") |
| 938 | "Wiltse and Stolterman [32]" | NordiCHI 2010 | OK |
| 938 | "Overbeeke and Wensveen [23] ... Gibson's ecological psychology (e.g. [15]) ... 'irresistables'" | DPPI 2003; Gibson 1979 | Numbers OK. Spelling slip: the text has "irresistables"; the source title and reference 23 both have "irresistibles" |
| 938 | "future practice [13, 14, 16, 30, 36]"; "Sengers and her colleagues [30]"; "practical problems' [13]" | manifestos | OK |
| 938 | "Forlizzi's Product Ecology [10]" | IJD 2008 | OK |
| 938 | "Djajadiningrat and his colleagues [6]" | PUC 2007 | OK |
| 938 to 939 | "Zimmerman & Forlizzi [34] argue ... 'agreed upon form of practice, evaluation, and outcome' ... 'extensible, systemic approaches'" | ref 34 is Zimmerman, Forlizzi and **Evenson** 2007 | **Mismatch: should be [33].** Both phrases are in Zimmerman and Forlizzi 2008 (CMU PDF, abstract: "without an agreed upon form of practice, evaluation, and outcome" and "extensible, systemic approaches to research through design are the most promising"); neither appears in the 2007 CHI paper (Berkeley PDF), which also has a third author |
| 939 | "Stolterman [31] ... ultimate particulars" | IJD 2008 | OK |
| 939 | "(see [5] p. 70)"; "(see [5] for a valuable overview)" | Curd and Cover | OK (p. 70 content not checked) |
| 939 | "Feynman [8] ... [8] p. 156"; "[8] p. 163" | Feynman | OK (page numbers not checked against a copy) |
| 939 | "'irrefutability is not a virtue ...' ([24], p. 7)"; "falsifiability ... [24]"; "([24] p 7)" | Popper, in Curd and Cover pp. 3 ff. | OK (p. 7 consistent with the reprint) |
| 939 | "According to Kuhn [17] ... 'puzzle-solving' ... ([17] p 67)" | Structure 1970 | Number OK; the exact wording and p. 67 could not be confirmed (Kuhn's known phrase is "failure to achieve a solution discredits only the scientist and not the theory") |
| 939 | "other accounts of theory, for instance from the humanities ... [2]" | humanities panel | OK |
| 939 | "Newton ... could not account for the moon's orbit [3]" | Collins 1994 | **Mismatch: should be [5]** (Curd and Cover), per the pitfall list; Collins is on social-science consensus |
| 939 to 940 | "Lakatos [19]"; "([19] p. 24)" twice | Lakatos, in Curd and Cover pp. 20 ff. | OK |
| 940 | "planet Neptune [5]" | Curd and Cover | OK |
| 940 | "'homo ludens' [13], 'suppleness' [16], ... [6]"; "'the self' [36]" | Gaver 2009, Isbister and Höök, Djajadiningrat, Zimmerman 2009 | OK |
| 940 | ambiguity block quote "[14]" | Gaver, Beaver, Benford 2003 | OK (number fits; quotation not checked against the full text) |
| 940 | "Popper's [24] criterion of falsifiablity" | Popper | OK ("falsifiablity" is a typo in the text) |
| 940 | "wicked problems [26]"; "(c.f. [7]."; "changes the context of its own activities [21]" | Rittel and Webber; Fallman; Law and Urry | OK (unclosed parenthesis is a typo) |
| 940 | "Zimmerman et al.'s [34] formulation, on making the 'right thing'" | Z, F and E 2007 | OK ("make the right thing" is in the 2007 paper) |
| 941 | "As Stolterman [31] puts it, 'designers can be prepared-for-action but not guided-in-action ...' (p 61)" | IJD 2008 | OK (phrase on p. 61) |
| 941 | "'Reflective Design' [30] ... 'Supple Interfaces' [16] ... 'Designing for the Self' [36]" | | OK |
| 941 | "Kuhn [17] in his seminal account" | | OK |
| 941 | "(see [29])" | Sengers 2006 | OK as a pointer (entry unverifiable) |
| 941 | "appeal for more agreement about the 'methods, approaches, and outcomes'" | no number | Quotation without a citation; the words come from the [9] quotation on p. 937 |
| 941 to 942 | "According to Kuhn [17], the beginning of a scientific study" | | OK |
| 942 | "test and produce theory [33, 35]"; "formalism and standardisation [9, 35]" | | OK |
| 942 | "as Kuhn [17] suggests"; "Latour [20]"; "Lakatos [19] and Kuhn [17, 18]" | | OK |
| 943 | "disagreements he witnessed amongst social scientists ([17]; c.f. [3])" | Kuhn 1970; Collins 1994 | OK (Collins fits the social-science point) |
| 943 | "(see [18] for one definition of 'better')" | Kuhn 1977 | OK |
| 943 | "(c.f. [20])" | Latour | OK |
| 943 | "Osborne and Rose [22]"; "Law and Urry [21]"; "[21] p. 397" | | OK (p. 397 inside 390-410) |
| 943 | "guiding design practice [33, 35]"; "([27] cited in [31])" | | OK |
| 944 | "(e.g. [6])"; "rotary controller for a microwave oven [6]" | Djajadiningrat 2007 | OK |
| 944 | "(c.f. [7])" | Fallman | OK |
| 944 | "As Stolterman [31] suggests ... ([31] p. 59)" | IJD 2008 | OK ("ultimate particular ... same dignity and importance as truth in science" on p. 59) |
| 944 | "As Carroll and Kellogg [3] have argued, a designed artifact is a 'theory nexus'" | Collins 1994 | **Mismatch: no entry exists.** Intended work: Carroll, J.M. and Kellogg, W.A. (1989). Artifact as theory-nexus: hermeneutics meets theory-based design. Proc. CHI '89, pp. 7-14 (https://api.crossref.org/works/10.1145/67449.67452; also ACM SIGCHI Bulletin 20(SI), 7-14) |
| 944 | "(see [12])" design space | Gaver 2011 | OK |
| 944 | www.designmuseum.org and www.vitsoe.com (p. 945) | inline URLs | Not in the list; acceptable as inline web references |
| 945 | "Alexander's [1] design patterns" | Alexander 1977 | OK |
| 946 | refs 18, 19, 24: "Reprinted in [3]" | Collins 1994 | **Mismatch: should be [5]** |

### Works cited but missing from the list
- Carroll and Kellogg 1989, "Artifact as theory-nexus: hermeneutics meets theory-based design", CHI '89, pp. 7-14 (cited as [3] on p. 944).

### List entries never cited in the text
- 25 (Pullin 2010).
- 28 (Schön 1983).

All other numbers 1 to 36 appear at least once in the body.

## Wrong references, in short

1. Ref 1: Murray Silverstein, the third principal author of A Pattern Language, is omitted.
2. Ref 5: published in New York, not London; "WW." should be "W.W."; Curd and Cover are editors.
3. Ref 7: "Fallmann" should be "Fallman".
4. Ref 15: "Haughton Mifflin" should be "Houghton Mifflin".
5. Ref 18: "Reprinted in [3]" should be "Reprinted in [5]".
6. Ref 19: "Reprinted in [3]" should be "Reprinted in [5]".
7. Ref 23: "Wesveen" should be "Wensveen".
8. Ref 24: "Reprinted in [3]" should be "Reprinted in [5]".
9. Ref 25: Design and Culture 2(3), not 2(1); title lacks its subtitle "from Social Mobiles to the Museum of Lost Interactions and Six Speaking Chairs".
10. Ref 26: title should read "Dilemmas in a General Theory of Planning"; original is Policy Sciences 4(2), 155-169, 1973 (the 1974 reprint venue is unconfirmed).
11. Ref 35: subtitle is "towards a formalization of a research approach", not "toward".
12. In-text p. 939: "[3]" for Newton and the moon should be [5].
13. In-text p. 939: "Zimmerman & Forlizzi [34]" with the quotes "agreed upon form of practice, evaluation, and outcome" and "extensible, systemic approaches" should be [33].
14. In-text p. 944: "Carroll and Kellogg [3]" points at Collins; the Carroll and Kellogg 1989 CHI paper is missing from the list.
15. In-text p. 938: "irresistables" should be "irresistibles" (as in ref 23 and the source title).
16. In-text p. 937: the [33] quotation reads "apply in research in practice"; the source has "apply in research and practice".
17. In-text p. 941: the quotation "'methods, approaches, and outcomes'" carries no citation (it is from [9]).
18. Refs 25 and 28 are listed but never cited.

## Unverifiable

- Ref 29, Sengers (2006), "Must design become 'scientific'?", DIS'06 workshop paper: the workshop exists, the paper was not found.
- Ref 26, the 1974 Design Research and Methods 8(1), 31-39 reprint venue (the title error is confirmed regardless).
- Ref 33, the series label "Human-Computer Interaction Institute, Paper 37" (the work and its 2008 date are confirmed via KiltHub; repository.cmu.edu refused the connection).
- Ref 8, a 1989 printing of Feynman's book (the book and MIT Press are confirmed; first published 1965).
- In-text [17] p 67 wording of the Kuhn quotation, [5] p. 70, [8] pp. 156 and 163, and the [14] block quotation: the citation numbers fit, but the page numbers and wording were not checked against copies.
