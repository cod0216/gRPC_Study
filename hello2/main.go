// go.mod는 모듈이고 간단히 build.gradle 이나 package.conf 이랑 같다.
// 라이브러리 디펜던시 관리해주는 모듈 파일이다. -> go.mod
//설명 https://darkstart.tistory.com/157

package main // go는 반드시 패키지로 시작 해야된다.
// 이 코드가 속한 패키지가 무엇인지 나타낸다.
// 패키지는 코드를 묶는 단위

// 다른 패키지 명들은 아무거나 적을 수 있지만
// main은 특별한 패키지 (아무렇게 적으면 안됨)
// main은 프로그램 시작점을 포함하는 패키지

// 어디서 부터 실행 할 것인가(프로그램 starting 지점)
// 시작 포인트는 하나여야됨, 그게 main패키지임

// go는 어떻게 구성되어 있나
// main 패키지 하나와 여러게의 다른 패키지로 구성될 수 있음


import "fmt" // 나는 fmt 패키지를 가져와 사용하겠다.
// import - 가저오다
// fmt - 패키지 명

// package main
// import main - 이렇게 자기 자신을 가져올 순 없음, 또한 import main 이렇게 다른 main을 가져올 수 도 없음

func main() { 
	// 이것은 주석입니다.

	// func 함수 라는 뜻, 즉 함수 선언
	// main() - 함수 이름

	// 함수 이름은 영문자로 시작하는거 아무거나 적으면 됨
	// 단 main 키워드는 특별한 키워드
	// 프로그롬 시작점 이라는 뜻
	
	// package main - 프로그램 시작점
	// 프로그램 시작점에 main 함수를 가지고 있는 애임

	// 프로그램이 로드 되면 main 함수부터 시작된다.

	// 즉 이 곳에서 { 여기! } 시작한다. 한 줄 씩!

	fmt.Println("hello go world") // fmt에 포함된 함수명


}