(use utils)
(use srfi-1)

;This solution in painfyully slow, but it works. (It has to go through 138 rounds)
(define (evalLine line)

	(if (equal? "+" (string-take line 1))
		(string->number (string-drop line 1))
		(* -1 (string->number (string-drop line 1)))
	)
	

)
(define (procLines lines acc visited loopCount)

	(display loopCount)
	(display ".")
	(if (member acc visited)
		acc	
		(if (null? lines)
			(procLines (read-lines "data.txt") acc visited (+ 1 loopCount))
			(procLines (cdr lines) (+ acc (evalLine (car lines))) (append (list acc) visited) loopCount)
		)
	)	
)

(define (main args)
	(display (procLines (list) 0 (list) 1))
)

