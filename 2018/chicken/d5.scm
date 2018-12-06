(use utils)
(use srfi-1)
(use srfi-13)
(use srfi-69)
(use vector-lib)


(define (build-remove-map line removedMap index)
    (if (= index (- (string-length line) 1 ))
        removedMap
        (begin

        )
    )
)
(define (main args)
    ;(define a (procLines 2 (read-lines "d5data.txt") 0 ))
    (define removedMap (build-remove-map "dabAcCaCBAcCcaDA" (make-hash-table) 0)
    (print (hash-table-keys removedMap))
	
)