import * as fs from 'fs';


function main() {
   const input_file = fs.readFileSync('day01.txt','utf-8') 

   let input = input_file.split(/\r?\n/).filter((line) => line);

   console.log(input[input.length-1]);

   console.log("Part 1 Result: ", part1(input));
   console.log("Part 2 Result: ", part2(input));
}

function part1(input: string[]) {
    let list1: number[] = []
    let list2: number[] = [];
    
     input.forEach((pair: string) => {
         let vals = pair.split('   ');
         list1.push(+vals[0]);
         list2.push(+vals[1]);
     });


     //console.log(Math.max(...list2));

     list1.sort();
     list2.sort();
     //console.log(`Size check: list1 size = ${list1.length}, list2.length = ${list2.length}`);

     //console.log('First slice of list1: ', list1.slice(-5, 5));
     //console.log('First slice of list2: ', list2.slice(-5, 5));

     let sum = 0;
     for (let i = 0; i < list1.length; i++) {
        //console.log(`Iterating for: ${i}`, list1[i], list2[i]);
        sum += Math.abs(list1[i] - list2[i]);
     }

    //console.log(`Sum is ${sum}`);
    return sum;
}


function part2(input: string[]) {
    let list1: number[] = []
    let list2: number[] = [];
    
    input.forEach((pair: string) => {
        let vals = pair.split('   ');
        list1.push(+vals[0]);
        list2.push(+vals[1]);
    });

    list1.sort();
    list2.sort();

    let countMap = new Map<number, number>();

    list1.forEach((val) => {
        countMap.set(val, 0);
    });

    list2.forEach((val) => {
        countMap.has(val) && countMap.set(val, countMap.get(val)! + 1);
    });


    let sum = 0;
    countMap.forEach((key, val) => {
        //console.log(key, val);
        sum += key * val;
    });
    return sum;
}


main();
