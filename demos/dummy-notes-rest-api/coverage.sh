#!/bin/bash

coverage_dir=".cover"
coverage_file="cover.out"
coverage_output="${coverage_dir}/${coverage_file}"
coverage_mode="count"
integration_build_tag="intgr"
help_message="coverage.sh
     \t--html      Additionally creates HTML report and open it in browser\n
     \t\t--${integration_build_tag}     Runs all tests including the ones with ${integration_build_tag} tag\n
     \t\t--help      Displays info on how to use the coverage script
"

generate_coverage_data() {
    build_tags=$1
    rm -rf ${coverage_dir}
    mkdir ${coverage_dir}
    if ! go test -cover -covermode="${coverage_mode}" -coverprofile="${coverage_output}" -coverpkg=./... ./... -tags="${build_tags}"
    then
        exit 1
    fi
}

show_coverage_report() {
    coverage_report_type=$1
    go tool cover -${coverage_report_type}="${coverage_output}"
}

run() {
    coverage_report_type="func"
    build_tags=""
    for i in "$@"
    do
        case "$i" in
            "")
            ;;
            --html)
                coverage_report_type="html"
            ;;
            --intgr)
                build_tags="${integration_build_tag}"
            ;;
            --help)
                echo -e ${help_message}
                exit 0
                ;;
            *)
            echo >&2 "error: invalid flag: $i. Check out: 'coverage.sh --help' for more info"; exit 1
            ;;
        esac
    done
    generate_coverage_data ${build_tags}
    show_coverage_report ${coverage_report_type}
}

run $@
