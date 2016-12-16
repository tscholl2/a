for numDice=1:10
    rolls = zeros(1000,1);
    for i=1:1000
        curMax = 0;
        curNumTens = 0;
        for die=1:numDice
            roll = randi(10,1,1);
            if roll == 10
                curNumTens = curNumTens + 1;
                curMax = 0;
            else
                if roll > curMax
                    curMax = roll;
                end
            end
        end
        rolls(i) = curNumTens*10+curMax;
    end
    subplot(10,1,numDice)
    histogram(rolls)
    axis([0 60 0 400])
end


for numDice=1:10
rolls = zeros(1000,1);
for i=1:1000
    curMin = 10;
    for die=1:numDice
        roll = randi(10,1,1);
        if roll < curMin
            curMin = roll;
        end
    end
    rolls(i) = curMin;
end
    subplot(10,1,numDice)
    histogram(rolls)
    axis([0 10 0 400])
end
