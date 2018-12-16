(use utils)
(use list-utils)
(use vector-lib)
(use srfi-1)
(use srfi-69)
(load "jsi-matrix.scm")

(define (cleanTokens line)
    (define t (string-tokenize (string-translate line "," )) )
    (list (string->number (first t)) (string->number(second t)))
)


(define (map-lines->cleanTokens lines acc)
	(if (null-list? lines)
		acc
		(map-lines->cleanTokens (cdr lines) (append acc (list(cleanTokens (car lines)))))
		)
)

(define (distance x1 y1 x2 y2)
    (+ (abs (- x1 x2)) (abs (- y1 y2)))
)

(define (calc-cell-closest x y cell data)
    (define grid (second data))
    (define points (first data))
    ;(print (list "cell" x y cell))
    (fold (lambda (point index) 
        ;(print (list "L" index point))
        (define cell (grid-ref grid x y))
        (define d (distance (first point) (second point) x y))
        (define c (append cell (list(list point d))))
        ;(print(list "C" c d))
        (define sorted (sort c less-point-distance))
        
        (grid-set! grid x y sorted)
    (+ index 1)) 0 points)

)

(define (calc-cell-total-distance x y cell data)
    (define grid (second data))
    (define points (first data))
    ;(print (list "cell" x y cell))
    (define acc (fold (lambda (point index) 
            ;(print (list "L" index point))
            (define cell (grid-ref grid x y))
            (define d (distance (first point) (second point) x y))
            ;(define c (append cell (list d)))
            
            ;(grid-set! grid x y c)
        (+ index d)) 0 points)
    )
        ;(print (list "acc" x y acc))
        (grid-set! grid x y acc)
)
(define (calc-cell-not-single x y cell grid)
    
    (define cell (grid-ref grid x y))
    (if (= (second(first cell)) (second(second cell))   )
        (grid-set! grid x y (list (list x y) -1))
        (grid-set! grid x y (first cell))
    )
)


(define (grid-fold grid f seed)
    (vector-fold (lambda (y seed row) 
                    (begin 
                        ;(print (list "row" y seed row))
                        (vector-fold (lambda (x seed cell) 
                            (begin 
                                (f x y cell seed)
                            )) seed row)
                    )) seed grid)
)


(define (count-less-than x y cell seed)
    (print (list "count-less-than" x y cell seed))
    (if (< cell 10000 )
        (+ 1 seed)
        seed
    )
)


(define (calc-area x y cell point-table)
    
    (define dist (second cell))
    (define key (->string (first cell) ))
    
    (define oldVal (hash-table-ref/default point-table key -1))
    (if (> oldVal -1)
        (hash-table-set! point-table key (+ 1 oldVal))
    )
)


(define (find-infinite-area! x y cell data)    
    (define boundaries (second data))
    ;(print (list "xxx" x y (first cell) boundaries ))

    (if
        (or
            (= x (first boundaries))
            (= x (second boundaries))
            (= y (third boundaries))
            (= y (fourth boundaries))
        )
        (begin
            ;(print (list "---" x y (first cell) ))
            (define point-table (first data))
            (define key (->string (first cell)))
            (hash-table-set! point-table key -1 )
        )
    )
)

(define (less-point-distance pa pb) 
    ;(print (list pa pb))
    (< (second pa) (second pb))
)
(define ((is-finite boundaries) point)
;(print boundaries)
;(print point)
    (not
        (or
            (= (first point) (first boundaries)) ;minX
            (= (first point) (second boundaries)) ;maxX
            (= (second point) (third boundaries)) ;minY
            (= (second point) (third boundaries)) ;maxY
        )
    )
)
(define (point->hash-pair point)
    (cons (->string point) 0)
)
(define (print-hash-pair key val)
    (print (list "# " key ":" val))
)
(define (main args)

    (define lines (read-lines "d6data.txt"))
    (define points (map-lines->cleanTokens lines (list)))
    (define minX (- (fold (lambda (point minX) (min (first point) minX) ) 100000 points) 1))
    (define minY (- (fold (lambda (point minY) (min (second point) minY) ) 100000 points) 1))
    (define maxX (fold (lambda (point maxX) (max (first point) maxX) ) 0 points) )
    (define maxY (fold (lambda (point maxY) (max (second point) maxY) ) 0 points) )
    (define boundaries (list minX maxX minY maxY))
    ;(print (list "bounds" boundaries))
    (define grid (make-grid (+ 1 maxX) (+ 1 maxY) (list)))
    
    (grid-for-each grid calc-cell-total-distance (list points grid))
    ;(print grid)
    ;(define sum 0)
    ;(grid-for-each grid count-less-than sum)
    (print (list "sum" (grid-fold grid count-less-than 0)))
 
    
    ;Part one >>
    ; (grid-for-each grid calc-cell-closest (list points grid))
    ; (grid-for-each grid calc-cell-not-single grid)
    
    ; ;(print grid)

    ; (define point-table (alist->hash-table (map point->hash-pair points)))
    ; ;(print "A")
    ; ;(hash-table-walk point-table print-hash-pair)

    ; (grid-for-each grid find-infinite-area! (list point-table boundaries))
    ; ;(print "B")
    ; ;(hash-table-walk point-table print-hash-pair)

    ; (grid-for-each grid calc-area point-table)
    ; ;(print "C")
    ; (hash-table-walk point-table print-hash-pair)
    ; ;<<<<< Part one
)