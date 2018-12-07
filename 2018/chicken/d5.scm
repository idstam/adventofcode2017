(use utils)
(require-extension list-utils)
(use srfi-1)
(use srfi-13)
(use srfi-69)
(use vector-lib)


(define (delete-n l n)
  (if (= n 0) 
      (cdr l)
      (append (list (car l)) (delete-n (cdr l) (- n 1))))
)


(define (poly-pair? a b)
    (print (list a b ) (= 32 (abs (- a b) )))
    (= 32 (abs (- a b) ))
)

(define (find-first-poly-pair int-list i)
    (if (length=1? int-list )
        #f
        (begin
            (define a (first int-list))
            (define b (second int-list))
            (if (poly-pair? a b)
                i
                (find-first-poly-pair (cdr int-list) (+ 1 i))
            )
        )
    )
)


(define (filter-poly-pairs int-list i)
    (if (= (length int-list) (+ 1 i))
        int-list
        (begin
            (define a (first int-list))
            (define b (second int-list))
            (if (poly-pair? a b)
                (begin 
                    (define a laksflkasdjfklasdjfas)asdfasdfasdfasdf
                    (define la (delete-n int-list a))
                    (define lb (delete-n la (+ 1 a)))
                    (filter-poly-pairs lb 0)
                )
                (filter-poly-pairs int-list (+ 1 i))
            )
        )
    )
)


(define (main args)
    (define int-list (map-in-order char->integer (string->list (car (read-lines "d5data.txt")))))
    (print int-list)
    (define filtered-list (filter-poly-pairs int-list 0))
    (print filtered-list)
    ;(print (find-first-poly-pair int-list 0))
    (print (list->string (map-in-order integer->char filtered-list)))

)