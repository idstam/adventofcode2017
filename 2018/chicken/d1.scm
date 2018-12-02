(use utils)
(use srfi-1)

(define (evalLine line)
	(if (equal? "+" (string-take line 1))
		(string->number (string-drop line 1))
		(* -1 (string->number (string-drop line 1)))
	)
	
)
(define (procLines lines acc)
	(if (null-list? lines)
		acc
		(procLines (cdr lines) (+ acc (evalLine (car lines))))
	)
)

(define (main args)
	(display (procLines (read-lines "data.txt") 0 ))
)

