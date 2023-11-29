#! bin/bash
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq --argjson my_hero "$HERO_ID" '.[] | select(.id == $my_hero) .connections.relatives' | sed 's/"//g'