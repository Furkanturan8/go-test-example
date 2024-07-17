# Örnek Test Fonksiyonları 

## 1. TestFuncName(*testing.T):

Bir test fonksiyonu tanımlar. *testing.T parametresi üzerinden test durumunu kontrol ederiz.


```golang
func TestAddition(t *testing.T) {
    result := Add(1, 2)
    expected := 3
    if result != expected {
        t.Errorf("Addition result was incorrect, got: %d, want: %d.", result, expected)
    }
}
```
## 2.  t.Errorf(format string, args ...interface{}):

Bir hata durumunda testi başarısız yapar ve hata mesajını gösterir. 

```golang
if result != expected {
    t.Errorf("Addition result was incorrect, got: %d, want: %d.", result, expected)
}
```

## 3.  t.Run(name string, f func(*testing.T)):

Test fonksiyonlarını gruplamak için kullanılır. Alt test grupları oluşturur.

```golang
func TestMathFunctions(t *testing.T) {
    t.Run("Addition", testAddition)
    t.Run("Subtraction", testSubtraction)
}
```

## 4. t.Skip(reason string):

Testin belirli bir nedenle atlanmasını sağlar.

```golang
func TestFeatureX(t *testing.T) {
    if !isFeatureXSupported() {
        t.Skip("Feature X is not supported in this environment.")
    }
    // Test Feature X functionality
}
```

## 5. t.Fatal(args ...interface{}):

Anında bir hata raporu oluşturarak testi başarısız yapar.

```golang
if err != nil {
    t.Fatal("Error occurred:", err)
}
```

## 6. t.Logf(format string, args ...interface{}):

Bilgi mesajları yazdırır, ancak test başarısını etkilemez.

```golang
t.Logf("Starting test for function %s.", functionName)
```

## 7. assert.Equal(t testing.T, expected, actual interface{}) (github.com/stretchr/testify/assert):

Beklenen ve gerçek sonuçları karşılaştırır ve eğer farklıysa testi başarısız yapar.

```golang
assert.Equal(t, 3, result)
```

## 8. assert.NotNil(t testing.T, object interface{}) (github.com/stretchr/testify/assert):

Bir nesnenin nil olup olmadığını kontrol eder.

```golang
assert.NotNil(t, result)
```

## 9. assert.NoError(t testing.T, err error) (github.com/stretchr/testify/assert):

Bir hata olmadığını kontrol eder.

```golang
assert.NoError(t, err)
```

## 10. assert.Contains(t testing.T, haystack interface{}, needle interface{}) (github.com/stretchr/testify/assert):

Bir koleksiyonun içinde belirli bir öğenin bulunup bulunmadığını kontrol eder.

```golang
assert.Contains(t, collection, item)
```

## 11. assert.Len(t testing.T, object interface{}, length int) (github.com/stretchr/testify/assert):

Bir koleksiyonun belirli bir uzunlukta olduğunu kontrol eder.

```golang
assert.Len(t, collection, expectedLength)
```

## 12. assert.ElementsMatch(t testing.T, expected, actual interface{}) (github.com/stretchr/testify/assert):

İki koleksiyonun elemanlarının eşleştiğini kontrol eder.

```golang
assert.ElementsMatch(t, expectedCollection, actualCollection)
```



