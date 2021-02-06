duration=1m
for rate in 1 10 100; do
    echo "Running at rate "$rate"...";
    ./loader -duration=$duration -rate=$rate -logfile=info$rate-$duration.log
done

