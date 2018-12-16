(use vector-lib)

(define (make-grid-row conf)
    ;(print (list "make-matrix-row" conf))
    (make-vector (first conf) (second conf))
)

(define (make-grid colCount rowCount fill)
    (define rows (make-vector rowCount 7))
    (vector-fold (lambda (y seed row) 
                (vector-set! rows y (make-vector colCount fill))
                0) 0 rows)
     rows
)

(define (grid-ref grid x y)
    (define row (vector-ref grid y))
    (vector-ref row x)
)
(define (grid-set! grid x y value)
    (define row (vector-ref grid y))
    (vector-set! row x value)
)

(define (grid-for-each grid f data)
    (vector-fold (lambda (y seed row) 
                    (begin 
                        ;(print (list "row" y seed row))
                        (vector-fold (lambda (x seed cell) 
                            (begin 
                                (f x y cell data)
                            0)) 0 row)
                    0)) 0 grid)
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

