#!/usr/bin/env bash
EXEC="./arvore_bench"
POOL="arvore-avl.goPool"
DUR=5 #segundos 
WORKLOADS=("insert" "update" "delete")

GO11MODULE=off ~/go-pmem/bin/go build -txn -o arvore_bench arvore_binaria_pmem.go

for W in "${WORKLOADS[@]}"; do
    echo "==============================================="
    echo "🚀 Rodando workload: $W por ${DUR}s"
    echo "==============================================="
    
    rm -f "$POOL"
    START=$(date +%s%N)
    timeout "$DUR" $EXEC -file "$POOL" -workload "$W" > tmp_$W.log 2>&1
    END=$(date +%s%N)

    tempo_decorrido_NS=$((END - START))
    tempo_decorrido_S=$(echo "scale=3; $tempo_decorrido_NS / 1000000000" | bc)

    COUNT=$(grep -Eo 'Inserções: [0-9]+' "tmp_${W}.log" | tail -1 | awk '{print $2}')
    if [ -z "$COUNT" ]; then
    COUNT=$(grep -Eo 'Ops: [0-9]+' "tmp_${W}.log" | tail -1 | awk '{print $2}')
    fi
    [ -z "$COUNT" ] && COUNT=0

    echo "⏱️ Tempo total: ${ELAPSED_S}s"
    echo "🧮 Operações detectadas: ${COUNT}"

    if [ "$COUNT" -gt 0 ]; then
      OPS_POR_S=$(echo "scale=3; $COUNT / $tempo_decorrido_S" | bc)
      echo "⚡ Taxa: $OPS_POR_S op/s"
    else
      echo "⚠️ Nenhuma operação detectada."
    fi
    echo
done
