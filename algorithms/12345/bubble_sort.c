#include<stdio.h>

#define LENGTH(array) ( (sizeof(array) ) / (sizeof(array[0])) )

#define swap(a,b)  (a^=b,b^=a,a^=b)

void bubble_sort(int a[], int n)
{
    int i, j;
    int flag;

    for(i = n-1; i > 0; i--)
    {
        flag = 0;
        
        for(j = 0; j < i; j++)
        {
            if(a[j] > a[j+1])
            {
                swap(a[j], a[j+1]);
                flag = 1;
            }
        }
        if (flag == 0)
            break;
    }

}

int main()
{
    int i;
    int a[] = {50,40,67,90,10,34,66,98,43};
    int ilen = LENGTH(a);

    printf("before sort:");
    for (i = 0; i < ilen; i++)
        printf("%d ", a[i]);
    printf("\n");

    bubble_sort(a, ilen);

    printf("agter sort:");
    for (i = 0; i < ilen; i++)
        printf("%d ", a[i]);
    printf("\n");
}



