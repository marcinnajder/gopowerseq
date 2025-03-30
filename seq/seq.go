package seq

// uwagi
// - warto pamietac ze pakiet wbudowany "slices." posiada juz kilka funkcji, aby zbednie konwertowac na Seq i powerseq
// - niby fajnie byloby pisac fluent interface jak ....Filter().Map() ale w go metody (funkcje zdefiniowane dla typow) nie moga byc generyczne wiec np
// funkcje Filter da sie napisac bo ma tylko jeden T, ale juz Map nie da sie napisac bo map T i R
// - z powyzszego powodu operatory piszemy jako zwykle funkcje, a nie metody, wtedy tez wygodnie kazdy moze dopisac swoj dodatkowy
// - z jednej strony chcemy moc wywolywac pojedyncze operatory, ale drugiej strony caly plus powerseq to mozliwosc wykorzystania kilku operatorow oraz ich leniwosci,
// kiedy zdecydujemy sie na prowadzenia Pipe(..., s => seq.Filter(s, ...) , s => ...) do skladania zapytania to pojawia sie wielka niewygoda polegajca na tym
// ze dla metod anonimowych zawsze i tak musimy pisac typy i kodu zaczyna sie pisac bardz duzo :( wiec tutaj moze operatory sa currowane wiec kod wyglada tak
// Pipe(..., seq.Filter(...), seq.Map(...))
// - gdy zdecydujemy sie na pisanie "curried funkcje" to pojaw sie niewygoda z pisania np takiego Take[string](10)(names) poniewaz tutaj inferencja typow nie dziala,
// trzeba jawnie podawac typ elementu w Seq[], wynika to z tego ze sama Take jest generyczna
// - nawet jak wprowadzilem sobie pomocnicze typu Operator, OperatorTR, ... to potem nie wszedzie mozna z nich skorzystac, albo sie nie znam albo po prostu go
// ma ograniczenia np .Find zwraca jakby "tuple" ale tuple to chyba nie jest typ danych a po prostu mozliwosc zwracania wielu rezultatu, takze .Concat przyjmowal
// zmienna ilosc argumentow zapisanych jako ...iter.Seq[T] i tam takze nie moglem zastosowac Operator[...,...]

// przyklad jak pisac wlasny modul ktory ma wiele pakierow
// - https://github.com/stretchr/testify, https://github.com/davecgh/go-spew
// - https://cgarciarosales97.medium.com/how-to-create-and-publish-packages-in-go-37482acfb4ab
// - https://softchris.github.io/golang-book/03-projects/03-create-shared-module/

// odpowiednik powerseq
// - https://github.com/ledongthuc/goterators tylko to dziala na slices i nie ma Pipe
// - https://serge-hulne.medium.com/iterators-map-filter-reduce-and-list-processing-in-go-golang-implementing-python-functional-2d24d780051f
// tutaj wykorzystuja channel do implemtacji operatorow :)
// -- https://github.com/serge-hulne/go_iter

// func Filter[T any](f func(T) bool, s iter.Seq[T]) iter.Seq[T] {
// 	return func(yield func(T) bool) {
// 		for v := range s {
// 			if f(v) {
// 				if !yield(v) {
// 					return
// 				}
// 			}
// 		}
// 	}
// }

// export type Func<T, R> = (arg: T) => R;
// export type Func2<T1, T2, R> = (arg1: T1, arg2: T2) => R;
// export type Func3<T1, T2, T3, R> = (arg1: T1, arg2: T2, arg3: T3) => R;

// export type Predicate<T> = (item: T, index: number) => boolean;
// export type Comparer<T> = (a: T, b: T) => number;
