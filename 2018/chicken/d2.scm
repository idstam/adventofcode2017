(use utils)
(use srfi-1)
(use srfi-69)

(define (letterCount letters letterTable)
	(if (null-list? letters)
		letterTable
		(begin
			(define val (hash-table-ref/default letterTable (car letters) 0))
			(hash-table-set! letterTable (car letters) (+ 1 val))
			(letterCount (cdr letters) letterTable)
		)
	)
)

(define (hasNum num line)
	(define letterTable  (letterCount (string->list line) (make-hash-table)))
	;(display (hash-table-ref letterTable (car (string->list line)) ))
	;(display ",")
	;(display num)
	(hash-table-fold  letterTable (lambda (key value foldedValue) 
		(if foldedValue
			#t
			(= value num)
		) ) #f)	
	
)

(define (procLines num lines acc )
	(if (null-list? lines)
		acc
		(if (hasNum num (car lines))
			(procLines num (cdr lines) (+ 1 acc))
			(procLines num (cdr lines) acc)
		)
	)	
)

(define (main args)
	(define a (procLines 2 (read-lines "d2data.txt") 0 ))
	
	(define b (procLines 3 (read-lines "d2data.txt") 0 ))
	(display (* a b))
	;(display (car (read-lines "d2data.txt")))
	;(display (car ))
	;(display (hasNum 2 (car (read-lines "d2data.txt"))))
)