#!/usr/bin/env bash -x

RED='\033[0;31m'
YELL='\033[1;33m'
NC='\033[0m' # No Color

# KEEP=""
# REPO=""

# Get all images at the given image repo
echo -e "${YELL}Getting all images${NC}"
echo "gcloud container images list --repository=${REPO} --format='get(name)'"
IMAGELIST=$(gcloud container images list --repository=${REPO} --format='get(name)')
if (( ${#IMAGELIST[@]} == 0 )); then
  echo -e "${RED}No image to cleanup${NC}"
  exit 0
fi
echo "$IMAGELIST"

while IFS= read -r IMAGENAME; do
  IMAGENAME=$(echo $IMAGENAME|tr -d '\r')
  echo -e "${YELL}Checking ${IMAGENAME} for cleanup requirements${NC}"

  # Get all the digests for the tag ordered by timestamp (oldest first)
  DIGESTLIST=$(gcloud container images list-tags ${IMAGENAME} --sort-by timestamp --format='get(digest)')
  DIGESTLISTCOUNT=$(echo "${DIGESTLIST}" | wc -l)

  if [ ${KEEP} -ge "${DIGESTLISTCOUNT}" ]; then
    echo -e "${YELL}Found ${DIGESTLISTCOUNT} digests, nothing to delete${NC}"
    continue
  fi

  # Filter the ordered list to remove the most recent 3
  DIGESTLISTTOREMOVE=$(echo "${DIGESTLIST}" | head -n -${KEEP})
  DIGESTLISTTOREMOVECOUNT=$(echo "${DIGESTLISTTOREMOVE}" | wc -l)

  echo -e "${YELL}Found ${DIGESTLISTCOUNT} digests, ${DIGESTLISTTOREMOVECOUNT} to delete${NC}"

  # Do deletion or say nothing to do
  if [ "${DIGESTLISTTOREMOVECOUNT}" -gt "0" ]; then
    echo -e "${YELL}Removing ${DIGESTLISTTOREMOVECOUNT} digests${NC}"
    while IFS= read -r LINE; do
      LINE=$(echo $LINE|tr -d '\r')
        gcloud container images delete ${IMAGENAME}@${LINE} --force-delete-tags --quiet
    done <<< "${DIGESTLISTTOREMOVE}"
  else
    echo -e "${YELL}No digests to remove${NC}"
  fi
done <<< "${IMAGELIST}"
