duration=5s
for rate in 100 500 1000 1500 2000 2500 3000; do
    echo "Running at rate "$rate"...";
    ./loader -duration=$duration -rate=$rate -logfile=info$rate-$duration.log
done

