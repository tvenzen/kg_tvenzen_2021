//Takes in command line arguments
var nums = process.argv.slice(2);

    //Assures arguments are converted to a string
    var num = nums.toString();

    //converts all numerals to words in the arry
    if (num.includes("0")){
        num = num.replace(/0/gi, "Zero");
    }
    if (num.includes("1")){
        num = num.replace(/1/gi, "One");
    }
    if (num.includes("2")){
        num = num.replace(/2/gi, "Two");
    }
    if (num.includes("3")){
        num = num.replace(/3/gi, "Three");
    }
    if (num.includes("4")){
        num = num.replace(/4/gi, "Four");
    }
    if (num.includes("5")){
        num = num.replace(/5/gi, "Five");
    }
    if (num.includes("6")){
        num = num.replace(/6/gi, "Six");
    }
    if (num.includes("7")){
        num = num.replace(/7/gi, "Seven");
    }
    if (num.includes("8")){
        num = num.replace(/8/gi, "Eight");
    }
    if (num.includes("9")){
        num = num.replace(/9/gi, "Nine");
    }
   

//displays output
process.stdout.write(num);