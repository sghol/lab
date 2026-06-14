import turtle

# Setup
screen = turtle.Screen()
screen.setup(600, 200)
screen.bgcolor("white")

pen = turtle.Turtle()
pen.speed(5)
pen.hideturtle()

# Draw a Python logo (simplified)
pen.penup()
pen.goto(-200, 30)
pen.pendown()
pen.color("blue", "lightblue")
pen.begin_fill()
for _ in range(36):
    pen.forward(10)
    pen.left(10)
pen.end_fill()

# Draw a simple turtle shape
pen.penup()
pen.goto(-120, 30)
pen.pendown()
pen.color("green")
pen.begin_fill()
for _ in range(3):
    pen.forward(30)
    pen.left(120)
pen.end_fill()

# Write the text
pen.penup()
pen.goto(-60, 0)
pen.pendown()
pen.color("black")
pen.write("Python Turtle Graphics", align="left", font=("Arial", 18, "bold"))

# Write the URL
pen.penup()
pen.goto(-100, -40)
pen.pendown()
pen.color("gray")
pen.write("Suherfe.blog.ir", align="left", font=("Arial", 12, "normal"))

screen.mainloop()
