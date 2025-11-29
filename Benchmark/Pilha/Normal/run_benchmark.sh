#!/usr/bin/env bash
EXEC="./pilha_bench"
DUR=15  # segundos
WORKLOADS=("insert" "read" "delete")

# Compilar com suporte a transações PMEM
GO111MODULE=off /home-ext/emilio/go-pmem/bin/go build -txn -o pilha_bench pilha.go

for W in "${WORKLOADS[@]}"; do
  echo "==============================================="
  echo "🚀 Rodando workload: $W por ${DUR}s"
  echo "==============================================="

  rm -f "$POOL"

  START=$(date +%s%N)
  timeout "$DUR" $EXEC -workload "$W" > tmp_$W.log 2>&1
  END=$(date +%s%N)

  ELAPSED_NS=$((END - START))
  ELAPSED_S=$(echo "scale=3; $ELAPSED_NS / 1000000000" | bc)

  COUNT=$(grep -Eo 'Empilhamentos: [0-9]+' "tmp_${W}.log" | tail -1 | awk '{print $2}')
  if [ -z "$COUNT" ]; then
    COUNT=$(grep -Eo 'Ops: [0-9]+' "tmp_${W}.log" | tail -1 | awk '{print $2}')
  fi
  [ -z "$COUNT" ] && COUNT=0

  echo "⏱️ Tempo total: ${ELAPSED_S}s"
  echo "🧮 Operações totais detectadas: ${COUNT}"

  if [ "$COUNT" -gt 0 ]; then
    OPS_PER_S=$(echo "scale=3; $COUNT / $ELAPSED_S" | bc)
    echo "⚡ Taxa: $OPS_PER_S op/s"
  else
    echo "⚠️ Nenhuma operação detectada."
  fi
  echo
done

