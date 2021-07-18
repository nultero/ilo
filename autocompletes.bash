_bx() {
    COMPREPLY=()
    local cur=${COMP_WORDS[COMP_CWORD]}

    local STARTOPTS=( "add" "edit" "help" "list" "remove" )
    local SECONDOPTS=( "alias" "event" "idea" "onetime" "recurrent" "todos" "wishlist" )
    local OPTS
    case $3 in 
    "add" | "edit" | "help" | "list" | "remove")
        OPTS=${SECONDOPTS[*]}
    ;;
    "event" | "idea" | "onetime" | "recurrent" | "todos" | "wishlist")
        OPTS=()
    ;;
    *)
        OPTS=${STARTOPTS[*]}
    ;;
    esac
    COMPREPLY=( `compgen -W "${OPTS[*]}" $cur` )
}

complete -F _bx bx