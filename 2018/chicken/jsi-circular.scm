(use utils)
(use list-utils)
(use vector-lib)
(use srfi-1)
(use srfi-69)


(define (make-circular)
    (define empty-list (list ))
    (let ((v (make-vector 3 0)))
            (vector-set! v 0 0)
            (vector-set! v 1 0)
            (vector-set! v 2 (list))
            v
        )

)


(define (circular-pos? circular)
    (vector-ref circular 0)
)
(define (circular-pos! circular pos)
    (vector-set! circular 0 pos)
)

(define (circular-length circular)
(vector-ref circular 1)
)
(define (circular-content circular)
    (vector-ref circular 2)
)

(define (circular-append! circular item)
    (vector-set! circular 1 (+ 1 (length (circular-content circular))))
    (vector-set! circular 2 (append (circular-content circular) (list item) ) )
    circular
)


(define (circular-move-right! circular steps)
    
    (let (
        (pos (+ steps (circular-pos? circular)))
        )
        (vector-set! circular 0 (circular-abs-pos? circular pos) )
    )   
)
(define (circular-move-left! circular steps)
    (let (
        (pos (+  (* -1 steps) (circular-pos? circular)))
        )
        (vector-set! circular 0 (circular-abs-pos? circular pos) )
    )   
)
(define (circular-abs-pos? circular pos)
    (modulo pos (circular-length circular))
)
(define (circular-at-pos? circular)

    (if (zero? (circular-pos? circular))
        (car (circular-content circular))
        (car (drop (circular-content circular) (circular-pos? circular)))
    )
)
(define (circular-insert-at! circular pos data)
    (let(
            (head (take (circular-content circular) pos))
            (tail (drop (circular-content circular) pos))
        )
        (begin 
            (vector-set! circular 2 (append head (list(list data)) tail ))
            (vector-set! circular 1 (length (circular-content circular)))
            circular
        )
    )    
)




(define (test expression expected )
    (define result expression)

    (if (equal? (->string result) (->string expected))
        (print "ok")
        (print (list "got" (->string result) "expected" (->string expected)))
    )
    
)
(define (circular-test)
 (define c (make-circular))
 (set! c (circular-append! c (list 'a )))
 (set! c (circular-append! c (list 'b)))
 (set! c (circular-append! c (list 'c)))
 (set! c (circular-append! c (list 'd)))
 (set! c (circular-append! c (list 'e)))

 (test (circular-length c) 5 )
 (test (circular-pos? c) 0)
 (test (circular-abs-pos? c 0) 0)
 (test (circular-abs-pos? c 3) 3)
 (test (circular-abs-pos? c -1) 4)
 (test (circular-abs-pos? c 5) 0)
 (test (circular-abs-pos? c 6) 1)

 (circular-move-right! c 1)
 (test (circular-pos? c) 1)

 (circular-move-left! c 2)
 (test (circular-pos? c) 4)

 (circular-move-right! c 2)
 (test (circular-pos? c) 1)
 (test (circular-at-pos? c) (list 'b) )


 (circular-pos! c 3)
 (test (circular-at-pos? c) (list 'd) )


 (circular-pos! c 0)
 (circular-insert-at! c 0 'x )
 (test (circular-at-pos? c) (list 'x) )
 (test (circular-length c) 6)


 (circular-insert-at! c 3 'y )
 (circular-pos! c 3)
 (test (circular-at-pos? c) (list 'y) )

;   (define c2 (make-circular))
;   (print c2)

;   (set! c2 (circular-append! c2 (list 'a )))
;   (print c2)
;   (circular-insert-at! c2 0 'x )
;   (print c2)

)