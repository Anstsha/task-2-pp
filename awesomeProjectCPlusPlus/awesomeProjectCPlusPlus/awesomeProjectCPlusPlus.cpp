#include <iostream>
#include <vector>
#include <omp.h>
#include <thread>
#include <numeric>

std::vector<int> generateArray(int size) {
    std::vector<int> data(size);
    for (int i = 0; i < size; ++i) {
        data[i] = i + 1;
    }
    return data;
}

long long sequentialSquareSum(const std::vector<int>& data) {
    long long result = 0;

    for (int i = 0; i < static_cast<int>(data.size()); ++i) {
        result += static_cast<long long>(data[i]) * data[i];
    }

    return result;
}

long long parallelSquareSum(const std::vector<int>& data) {
    long long result = 0;

#pragma omp parallel for reduction(+:result)
    for (int i = 0; i < static_cast<int>(data.size()); ++i) {
        result += static_cast<long long>(data[i]) * data[i];
    }

    return result;
}

void RunExample1() {

    std::vector<int> array = generateArray(100000000);

    double start = omp_get_wtime();
    long long result = sequentialSquareSum(array);
    double end = omp_get_wtime();

    std::cout << "Sum of array square elements: " << result << std::endl;
    std::cout << "Time without paralelism: " << end - start << std::endl;

    start = omp_get_wtime();
    long long resultParalel = parallelSquareSum(array);
    end = omp_get_wtime();
    std::cout << "Sum of array square elements: " << resultParalel << std::endl;
    std::cout << "Time with paralelism: " << end - start << std::endl;
}

void handleRequestSequential(int requestId) {
    std::cout << "Processing request " << requestId << "...\n";
    std::this_thread::sleep_for(std::chrono::seconds(2));
    std::cout << "Request  " << requestId << " processed\n";
}


void handleRequestParallel(int requestId) {
#pragma omp critical
    std::cout << "Processing request " << requestId << "...\n";

    std::this_thread::sleep_for(std::chrono::seconds(2));

#pragma omp critical
    std::cout << "Request " << requestId << " processed\n";
}

void RunExample2() {
    std::vector<int> requests = { 1, 2, 3, 4, 5 };

    std::cout << "Execution without parallelism:\n";
    double start = omp_get_wtime();

    for (int req : requests) {
        handleRequestSequential(req);
    }
    double end = omp_get_wtime();
    std::cout << "Execution time without parallelism: " << end - start << " ms\n";


    std::cout << "\nExecution with parallelism:\n";
    start = omp_get_wtime();
#pragma omp parallel for
for (int i = 0; i < static_cast<int>(requests.size()); ++i) {
    handleRequestParallel(requests[i]);
}
    end = omp_get_wtime();
    std::cout << "Execution time with parallelism: " << end - start << " ms\n";
}


struct Result {
    std::string Query;
    std::string Payload;
};

std::string simulateAPIRequest(const std::string& query) {
    std::this_thread::sleep_for(std::chrono::milliseconds(100));
    return "Result for " + query;
}

void printResults(const std::vector<Result>& results) {
    for (const auto& res : results) {
        std::cout << "Query: " << res.Query << ", Result: " << res.Payload << std::endl;
    }
}

std::vector<Result> processQueriesWithoutParallelism(const std::vector<std::string>& queries) {
    std::vector<Result> results;

    for (const auto& query : queries) {
        std::string payload = simulateAPIRequest(query);
        results.push_back({ query, payload });
    }

    return results;
}

std::vector<Result> processQueriesWithParallelism(const std::vector<std::string>& queries) {
    std::vector<Result> results;
    omp_lock_t lock;
    omp_init_lock(&lock);

#pragma omp parallel for
    for (int i = 0; i < queries.size(); ++i) {
        std::string query = queries[i];
        std::string payload = simulateAPIRequest(query);

        omp_set_lock(&lock);
        results.push_back({ query, payload });
        omp_unset_lock(&lock);
    }

    omp_destroy_lock(&lock);

    return results;
}

void RunExample3() {
    std::vector<std::string> queries = { "query 1", "query 2", "query 3", "query 4", "query 5" };

    double start = omp_get_wtime();
    auto resultsWithoutParallelism = processQueriesWithoutParallelism(queries);
    double end = omp_get_wtime();


    std::cout << "Execution without parallelism:\n";
    printResults(resultsWithoutParallelism);
    std::cout << "Time without parallelism: " << end - start << " ms\n";

    start = omp_get_wtime();

    auto resultsWithParallelism = processQueriesWithParallelism(queries);
    end = omp_get_wtime();

    std::cout << "\nExecution with parallelism:\n";
    printResults(resultsWithParallelism);
    std::cout << "Time with parallelism: " << end - start << " ms\n";
}

int main() {

#ifdef _OPENMP
    printf("_OPENMP Defined\n");
#else
    printf("_OPENMP UnDefined\n");
#endif

    RunExample1();
    RunExample2();
    RunExample3();
    return 0;
}