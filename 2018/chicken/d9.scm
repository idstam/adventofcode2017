(use utils)
(use list-utils)
(use vector-lib)
(use srfi-1)
(use srfi-69)
(load "jsi-circular.scm")

(define (round player scores marbles circle )
    ;(print (list "round" player scores marbles circle))
    (print (list "circle" circle))
    (if (null-list? marbles)
        888
        (begin
            (let (
                    (marble  (car marbles))
                    (pos (circular-abs-pos? circle (+ 2 (circular-pos? circle))))
                )
                
                    ; (print (list "pos" pos))
                    ; (print (list "circle" circle))
                    ; (print (list "marble" marble ))
                    (circular-insert-at! circle pos marble)
                    (round (+ 1 player) scores (cdr marbles) circle)            
                
            )
        )
    )
)

(define (main args)


(define player-count 9)
(let (
    (marbles (cdr (list-tabulate 26 values)))
    (circle (circular-append! (make-circular) 0 ))
    (scores (make-vector player-count 0 ))
)
    (print (round 0 scores marbles circle))
    ;(print circle)
    ;(print marbles)
)

;(circular-test) 




)