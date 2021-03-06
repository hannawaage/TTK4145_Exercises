with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;

procedure exercise8 is

    Count_Failed    : exception;    -- Exception to be raised when counting fails
    Gen             : Generator;    -- Random number generator

    protected type Transaction_Manager (N : Positive) is
        entry Finished;
        entry Wait_Until_Aborted;
        procedure Signal_Abort;
    private
        Finished_Gate_Open  : Boolean := False;
        Aborted             : Boolean := False;
    end Transaction_Manager;
    protected body Transaction_Manager is
        entry Finished when Finished_Gate_Open or Finished'Count = N is
        begin
	   
	   ------------------------------------------
	   -- PART 3: Modify the Finished entry
	   ------------------------------------------
        Finished_Gate_Open := Finished'Count /= 0;
            if not Finished_Gate_Open then
                Aborted := False;
            end if;
       ----- 

	   --------
        end Finished;

        procedure Signal_Abort is
        begin
            Aborted := True;
        end Signal_Abort;

        

	

	------------------------------------------
	-- PART 2: Create the Wait_Until_Aborted entry
	------------------------------------------
    entry Wait_Until_Aborted when Aborted is
        begin 
            if Wait_Until_Aborted'Count = 0 then
                Aborted := False;
            end if;
    end; 

    end Transaction_Manager;



    
    function Unreliable_Slow_Add (x : Integer) return Integer is
    Error_Rate : Constant := 0.15;  -- (between 0 and 1)
    begin
       if Random(Gen) > Error_Rate then
	  delay Duration(Random(Gen) * 5.0);
	  return X + 10;
       else 
	  delay Duration(Random(Gen) * 1.0);
	  raise Count_Failed;
       end if;
    end Unreliable_Slow_Add;

    function Add_Five (x : Integer) return Integer is
    begin
	  return X + 5;
    end Add_Five;



    task type Transaction_Worker (Initial : Integer; Manager : access Transaction_Manager);
    task body Transaction_Worker is
        Num         : Integer   := Initial;
        Prev        : Integer   := Num;
        Round_Num   : Integer   := 0;
    begin
        Put_Line ("Worker" & Integer'Image(Initial) & " started");

        loop
            Put_Line ("Worker" & Integer'Image(Initial) & " started round" & Integer'Image(Round_Num));
            Round_Num := Round_Num + 1;

            ------------------------------------------
            -- PART 1: Select-Then-Abort statement
            ------------------------------------------
            Prev := Num; 
            select
                -- eg. X.Entry_Call;
                -- code that is run when the triggering_alternative has triggered
                --   (forward ER code goes here)
                Manager.Wait_Until_Aborted;
                Num := Add_Five(Num);
                Put_Line ("  Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));
            then abort
                 begin
                 Num := Unreliable_Slow_Add(Num);
                 exception
                 when Count_Failed =>
                 Manager.Signal_Abort;
                 end;
                 Put_Line ("  Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));
                 Manager.Finished;
                -- code that is run when nothing has triggered
                --   (main functionality)
            end select;

                Prev := Num;
                delay 0.5;

        end loop;
    end Transaction_Worker;

    Manager : aliased Transaction_Manager (3);

    Worker_1 : Transaction_Worker (0, Manager'Access);
    Worker_2 : Transaction_Worker (1, Manager'Access);
    Worker_3 : Transaction_Worker (2, Manager'Access);

begin
    Reset(Gen); -- Seed the random number generator
end exercise8;

