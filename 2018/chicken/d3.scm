(use utils)
(use srfi-1)
(use srfi-69)

(define (cleanTokens lines)
	(display (car lines))
	(define cleanString (string-delete '@' (car lines)))
	(display cleanString)
)



(define (main args)
	(define a (cleanTokens (read-lines "d3data.txt")))
	
	
	(display a)
	;(display (car (read-lines "d2data.txt")))
	;(display (car ))
	;(display (hasNum 2 (car (read-lines "d2data.txt"))))
)