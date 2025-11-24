#!/usr/bin/env bash
#Benchmark para lista ligada persistente (insert-only e insert+update)
#Mede operações por tempo (op/ns e op/s)

EXEC="./arvore_bench"
DUR=5
WORKLOADS=("insert" "update" "delete")

#Compilar (use o compilador persistente)
GO11MODULE=off /home-ext/hreuter/go-pmem/bin/go build -txn -o arvore_bench arvore_binaria.go

for W in "${WORKLOADS[@]}"; do
  echo "=========================================="
  echo "🚀 Rodando workload: $W por ${DUR}s"
  echo "=========================================="

  START=$(date +%s%N)
  timeout "$DUR" $EXEC -workload "$W" >tmp_$W.log 2>&1
  END=$(date +%s%N)


  TEMPO_DECORRIDO_NS=$((END - START))
  TEMPO_DECORRIDO_S=$(echo "scale=3; $TEMPO_DECORRIDO_NS / 1000000000" | bc)

  # Extrai o número total de operações da saída
  CONT=$(grep -Eo "Inserções até agora: [0-9]+" tmp_$W.log | tail -1 | awk '{print $4}')
  if [ -z "$CONT" ]; then
    CONT=$(grep -Eo "Ops: [0-9]+" tmp_$W.log | tail -1 | awk '{print $2}')
  fi
  if [[ -z "$CONT" ]]; then
    CONT=0
  fi

  echo "⏱️ Tempo total: ${TEMPO_DECORRIDO_S}"
  echo "🧮 Operações totais detectadas: ${CONT}"

  if [ "$CONT" -gt 0 ]; then
    OPS_PER_S=$(echo "scale=3; $CONT/$TEMPO_DECORRIDO_S" | bc)
    echo "Taxa: $OPS_PER_S op/s"
  else
    echo "Nenhuma operação detectada"
  fi
  echo
done
