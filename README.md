# Autodiff (GO lang implementation)

You've probably heard about backpropagation, the very algorithm that people talk about constantly when they talk about Neural Networks. Well, people don't usually talk about Neural Nets. I hang out with really weird people who care about very weird things.

However, I wasn't so interested in the implementation of backpropagation that led me to this project that I've somewhat completed. I was reading Statistical Rethinking by McElreath *(great book by the way)* and he discussed about how the estimation of Bayesian models using the Hamiltonian method relies quite a bit on autodiff (or Auto-differentiation). Basically, when you have a nice differentiable function. which itself can be composition of functions and/or have usualy algebraic operations such as addition and multiplication, one can use the techniques we learn from Calculus to "automatically differentiate" the function given the input *x*. And *x* can be a vector, in which case, you would return a vector of its partial derivatives.

## How does it work?
I am no expert in this given how much I despise Calculus and Analysis as objects of mathematical study, but one can simply apply the usual Chain Rule for composition and Multiplication Rule for well, multiplication. One can apply the Division Rule as well, but one can choose not to be a contrarian and just use the Multiplication Rule. 

What about addition and subtraction? The good thing is that differentiation is a linear operation, and so the derivative of the sum is the sum of the derivatives! Yay!

## But SO WHAT if we can use Calculus Tricks?

The cool part of it all, is that once we know the derivative of each component, we can use the aforementioned rules, to decompose the functions into smaller functions. We can construct a **Computational Graph** that determines what has to be evaluated first etc. and calculate the derivative.

Like backprop, a huge part of it is to construct the computation graph. Unlike backprop, writing code for arbitrary functions is more difficult than I thought! Largely because backprop has the nice assumption that you are dealing with Linear + ReLU, and the application of Chain Rule is more straightforward from there. But that is just my assumption. I bet it is probably as hard or even worse.

I hope those folks at PyTorch are paid tonnes to this. I am not sure how one can create a solution that is comprehensive without blowing their heads off (as a solo dev of course). Unless they are unemployed like I am and are 10x, in which case, it might just be sad that they are unemployed.

# Reference

I used the resources below to help me understand the nuts and bolts of the algorithm. Though, I am not sure why their code works. (i.e. I don't think their calculation of the derivative from the computation graph is correct, but I am sure it is cuz they produced the right answer!)

https://www.csie.ntu.edu.tw/~cjlin/papers/autodiff/autodiff.pdf
https://github.com/ntumlgroup/simpleautodiff

# Limitations

Be warned, this is not a good implementation! These are a few caveats:

1. I did not apply multiplication rule, it is just way too complicated! So, don't use this if the function wrt to a variable $x_1$ is of the form $$f(x_1)g(x_1)$$
1. Because of (1), also don't use functions of the form $$\frac{f(x_1)}{g(x_1)}$$
1. No exhaustive testing! So use at your own risk!
