import manim as m

# How many seconds to write a single character?
WRITE_SPEED = 1/10


class MetricLoopFunction(m.Scene):
    """Metric loop function scene."""
    # pylint: disable=too-few-public-methods,invalid-name

    def construct(self):
        """Construct the scene."""
        # In this video, I'll try to describe a function that takes a specific
        # kind of *loop* as its input, and maps it to a real number.

        f = m.Tex("f").scale(2)
        unction = m.Tex("unction").scale(2).next_to(f, buff=0)
        word = m.VGroup(f, unction).center()
        self.play(m.Write(word))
        self.wait(1)

        fm = m.MathTex("f").scale(2).move_to(f)
        self.play(m.AnimationGroup(
            m.ReplacementTransform(f, fm),
            m.FadeOut(unction),
        ))
        self.wait(1)

        f = fm
        del fm

        # TODO: fix the kerning (set buf to the right value)!
        loop = m.Tex("(", "loop", ")", "$f(x)$").scale(2).next_to(f)
        self.play(m.Write(loop))

        self.wait(5)
        return


        #f = m.Tex("$f$").scale(2).next_to(function)
        line = m.VGroup(function, f).center()
        #self.play(m.AnimationGroup(
        #    m.Write(f),#, run_time=WRITE_SPEED),
        #    m.Write(unction)#, run_time=len("unction")*WRITE_SPEED),
        #))
        self.play(m.Write(line))
        self.wait(1)

        #self.wait(5)
        #return

        self.play(m.FadeOut(function))
        self.wait(1)

        # TODO:
        tex = m.MathTex(
            ":",
            #r"\bigcirc",
            "(loop)",
            "=",
            "r",
            ",",
            "r",
            r"\in",
            r"\mathbb{R}",
        ).scale(2).next_to(f)
        line.add(tex)

        self.play(m.Write(tex))

        # TODO.

        # If you don't know what a loop is, it's basically what it sounds like.

        # TODO.

        # You can think of it as a closed line, like a circle or some polygon,
        # but more precisely it is a *continuous function* that maps the *unit
        # interval* to some *topological space*, in such a way that the
        # boundaries 0 and 1 end up being mapped to the same value.

        # TODO.

        # The specific kind of loops we're interested in are the ones that map
        # this unit interval to a *metric space*. Let's call these *metric
        # loops*.

        # TODO.

        # The reason we want a metric space is because these are spaces in
        # which the *distance* between two points is precisely defined.

        # TODO.

        # But let's not go into more details about spaces and topology.
        # Instead, we'll use the good old 2D Euclidean space for our examples,
        # and maybe generalise for three or more dimensions later on.

        # TODO.

        self.wait(5)
        return

        # Imagine a circle,
        R = 2

        circle = m.Circle(radius=R, stroke_color=m.WHITE)
        self.play(m.Create(circle))
        self.wait(1/2)

        # with some origin point P0 and a direction along the circumference.
        dot_0 = m.Dot(point=circle.point_at_angle(0))
        dot_0_text = m.Tex("(1,0)").next_to(dot_0, m.RIGHT)
        self.play(m.AnimationGroup(
            m.FadeToColor(circle, m.GREY),
            m.GrowFromCenter(dot_0),
            m.FadeIn(dot_0_text, shift=m.LEFT/4),
            run_time=1/4,
            lag_ratio=0,
        ))
        self.wait(1/2)

        #dot_0_arrow = m.Vector(m.UP).shift(m.RIGHT*R)
        #self.play(m.Create(dot_0_arrow, run_time=1/4))

        dot_0_arc = m.Arc(
            radius=circle.radius,
            angle=m.PI/6,
        )
        # Create a tip on a separate, 0-width tangent line; see:
        # https://www.reddit.com/r/manim/comments/hqaxk7/curvedarrow_is_slightly_off
        dot_0_tip = m.Arc(
            radius=circle.radius,
            angle=m.PI/6# + 1/50,
        ).create_tip()
        #arrow_offset_fix = 1/250
        #dot_0_tan = m.TangentLine(circle, 1/12+arrow_offset_fix, length=2)
        #dot_0_tan = m.Line(
        #    start=dot_0_tan.start,
        #    end=dot_0_tan.end,
        #).set_stroke_width(0).add_tip()

        #dot_0_arrow = m.Group(dot_0_arc, dot_0_tan)
        self.play(m.AnimationGroup(
            m.Create(dot_0_arc),
            m.Create(dot_0_tip),
            #m.Create(dot_0_tip, run_time=1/4),
            #m.Create(tl),
            run_time=1/4,
        ))

        #self.add(dot_0)

        # We can use the distance from this origin point to define each point
        # on the circle.

        # Now, take some point P on the circle.
        #dot_1 = m.Dot(point=circle.point_at_angle(m.TAU/3))
        #self.add(dot_1)

        # We'll represent this point by the arc length from some origin point
        # P_0.

        self.wait(5)
