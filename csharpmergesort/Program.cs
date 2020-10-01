using System;
using System.Linq;
using GoFlag;
using System.Diagnostics;

namespace mergesort
{
    class Program
    {
        static void Main(string[] args)
        {
            var size = Flag.Uint("size", 0, "size of the array to generate");
            Flag.Parse();
            if (size == 0) {
                Flag.PrintDefaults();
                System.Environment.Exit(0);
            }

            var arr = random(size.Value);
            Console.WriteLine($"size: {arr.Length}");

            var sw = new Stopwatch();
            sw.Start();
            arr = MergeSort(arr);
            sw.Stop();
            var swFormat = "mm\\:ss\\.fffff";
            if (check(arr)) {
                Console.WriteLine($"sorted: {sw.Elapsed.ToString(swFormat)}");
            } else {
                Console.WriteLine($"nope...: {sw.Elapsed.ToString(swFormat)}");
            }
        }

        static bool check(int[] arr){
            if (arr.Length == 0) {
                return true;
            }
            var v = arr[0];
            for (var i = 1; i < arr.Length; i++) {
                if (v > arr[i]) {
                    return false;
                }
            }
            return true;
        }

        static int[] random(uint size) {
            Random r = new Random(DateTime.Now.Millisecond);
            var arr = new int[size];
            for (var i = 0; i < size; i++) {
                var v = r.Next();
                arr[i] = v;
            }
            return arr;
        }

        static int[] MergeSort(int[] arr) {
            if (arr.Length > 1) {
                int mid = (int)Math.Floor(arr.Length/2M);
                var a = MergeSort(arr.Take(mid).ToArray());
                var b = MergeSort(arr.Skip(mid).ToArray());
                //merge results
                var aIdx = 0;
                var bIdx = 0;
                for (int i = 0; i < arr.Length; i++) {
                    if (aIdx == a.Length) {
                        //we've exhausted array a, fill from b
                        arr[i] = b[bIdx++];
                    } else if (bIdx == b.Length) {
                        //we've exhausted array, fill from a
                        arr[i] = a[aIdx++];
                    } else if (a[aIdx] < b[bIdx]) {
                        //take from a and increment index (inline postfix operator)
                        arr[i] = a[aIdx++];
                    } else {
                        //take from b and increment index (inline postfix operator)
                        arr[i] = b[bIdx++];
                    }
                }
            }
            return arr;
        }
    }
}
