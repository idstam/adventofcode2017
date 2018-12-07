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

(define (nth l n)
  (if (or (> n (length l)) (< n 0))
    (error "Index out of bounds.")
    (if (eq? n 0)
      (car l)
      (nth (cdr l) (- n 1)))))

(define (poly-pair? a b)
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
            (define a (nth int-list i))
            (define b (nth int-list (+ 1 i)))
            (if (poly-pair? a b)
                (begin 
                    (define la (delete-n int-list i))
                    (define lb (delete-n la i)) ;Eftersom jag tog bort en precis är det samma index som ska tas bort igen
                    (filter-poly-pairs lb 0)
                )
                (filter-poly-pairs int-list (+ 1 i))
            )
        )
    )
)


(define (main args)
    (define int-list (map-in-order char->integer (string->list (car (read-lines "d5data.txt")))))
    (define filtered-list (filter-poly-pairs int-list 0))
    (define polymer (list->string (map-in-order integer->char filtered-list)))
    (print polymer)
    (print (string-length polymer))
)