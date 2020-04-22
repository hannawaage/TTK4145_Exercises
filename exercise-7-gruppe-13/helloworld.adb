with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;

procedure helloworld is 
    Count_Failed    : exception;    -- Exception to be raised when counting fails
    Gen             : Generator;    -- Random number generator

    function Unreliable_Slow_Add (x : Integer) return Integer is
    Error_Rate : Constant := 0.15;  -- (between 0 and 1)
    rand_num : Standard.Float := Random(Gen);
    c : Float := 2.5;
    d : Float := c*rand_num;
    a : Integer := x;
    begin
        if Error_Rate < rand_num then 
            delay Duration(d);
            a := x + 10;
            return a;
        else
            raise Count_Failed;
        end if;
    end Unreliable_Slow_Add;

    x : integer := 5; 
    a : integer := 0;
    begin
    a := Unreliable_Slow_Add(x);
    Put_Line(integer'Image(a));



end helloworld;