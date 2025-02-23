public class Prime{
    public static void main(String[] args) {
       Prime p1 =  new Prime();
       p1. primeNumber(50);
    }

    void primeNumber(int n){
        for(int i=1;i<=n;i++){
            if(i == 1 ){
                continue;
            }

             int count = 0;
            for(int j= 2;j<i;j++){
              if(i%j == 0){
                     count++;
              }
            }
            if(count == 0){
                System.out.println(i);
            }
        }
    }
}

// 2,3,5,7,11 
// = divide by 1 or itself only -> not devide by anyone other 