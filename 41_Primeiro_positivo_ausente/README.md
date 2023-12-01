**Pergunta:**

[LeetCode - First Missing Positive](https://leetcode.com/problems/first-missing-positive/)

```
Dado um array inteiro não classificado `nums`, retorne o menor número inteiro positivo ausente.
Você deve implementar um algoritmo que seja executado no tempo O(n) e use espaço auxiliar O(1).

**Exemplo 1:**
Entrada: nums = [1,2,0]
Saída: 3
Explicação: Os números no intervalo [1,2] estão todos na matriz.

**Exemplo 2:**
Entrada: nums = [3,4,-1,1]
Saída: 2
Explicação: 1 está na matriz, mas 2 está faltando

**Exemplo 3:**
Entrada: nums = [7,8,9,11,12]
Saída: 1
Explicação: O menor número inteiro positivo 1 está faltando.

**Restrições:**
- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1
```

<hr>

**Solução:**

Aqui está uma explicação passo a passo:


**Criação do Slice `index`:**

```go
index := make([]bool, len(nums)+1)
```

Cria um novo slice index de tipo bool com um comprimento de len(nums)+1. Este slice será usado para rastrear a presença de números positivos em nums.

Preenchimento do Slice index:

```go
for _, v := range nums {
    if v > 0 && v <= len(nums) {
        index[v] = true
    }
}
```

Percorre cada elemento v em nums. Se v é um número positivo e está dentro do intervalo válido, o índice correspondente em index é marcado como true, indicando que o número v está presente.

Encontrando o Menor Número Positivo Ausente:

```go
for i := 1; i <= len(nums); i++ {
    if index[i] == false {
        return i
    }
}
```

Itera sobre os números de 1 até o comprimento de nums. Retorna o primeiro número ausente, identificado pelo índice em index sendo false.

Se Todos os Números de 1 até len(nums) Estão Presentes:

```go
return len(nums) + 1
```

Se todos os números de 1 até o comprimento de nums estiverem presentes, o loop anterior não encontrará um número ausente. Nesse caso, return len(nums) + 1 é executado, indicando que o menor número positivo ausente é len(nums) + 1.