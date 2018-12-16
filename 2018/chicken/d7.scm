(use utils)
(use list-utils)
(use vector-lib)
(use srfi-1)
(use srfi-69)
(load "jsi-matrix.scm")



(define (cleanTokens line)
    (define t (string-tokenize line)) 
    (list (second t) (eighth t))
)

(define (lines->tokens lines acc)
	(if (null-list? lines)
		acc
		(lines->tokens (cdr lines) (append acc (list(cleanTokens (car lines)))))
		)
)


(define (print-hash-pair key val)
    (print (list "# " key ":" val))
)

(define (string-less? s1 s2)
	(<  (string-compare3 s1 s2) 0)
)
(define (string-eq? s1 s2)
	(=  (string-compare3 s1 s2) 0)
)


(define (find-first-done letter-table keys)
        ;(print (list "ffd" keys))
        (if (null-list? keys)
            "!"
            (begin
                (define sorted-keys (sort keys string-less? ))
                (define work (hash-table-ref letter-table (first sorted-keys) ))
                ;(print (list "first work" (first sorted-keys) work ))
                (if (null-list? work)
                    (first sorted-keys)
                    (find-first-done letter-table (cdr sorted-keys))
                )
            )
        )
)
(define (remove-work letter-table keys to-remove)
    (if (null-list? keys)
        (list)
        (begin
            (define all-work (hash-table-ref letter-table (first keys)))
            (define new-work (remove (lambda (w) (string-eq? w to-remove) ) all-work ))
            ;(print (list "remove-work" all-work new-work))
            (hash-table-set! letter-table (first keys) new-work)
            (remove-work letter-table (cdr keys) to-remove)
        )
    )
)

(define (do-work letter-table)
    (if (= 0 (hash-table-size letter-table))
        '()
        (begin
         (define work (find-first-done letter-table (hash-table-keys letter-table)))   
            (print work)
            (remove-work letter-table (hash-table-keys letter-table) work)
            (hash-table-delete! letter-table work)
            (do-work letter-table )
        )
    )
)

(define (main args)
    (define lines  (read-lines "d7data.txt") ) 
    (define tokens (lines->tokens lines (list)))

    (define letter-table (make-hash-table))
    (fold (lambda (token letter-table)
        (hash-table-set! letter-table (first token) (list) )
        (hash-table-set! letter-table (second token) (list) )
        letter-table
    ) letter-table tokens )

    ;(print (hash-table-keys letter-table))

    (fold (lambda (token letter-table)
        (define foo (hash-table-ref letter-table (second token) ))
        (define bar (append foo (list (first token)) ))
        (hash-table-set! letter-table (second token) bar )
        letter-table
    ) letter-table tokens )
    (hash-table-walk letter-table print-hash-pair)
    
    ;(find-first-done letter-table (hash-table-keys letter-table))
    (do-work letter-table)

)