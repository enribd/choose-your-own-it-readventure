FROM squidfunk/mkdocs-material

RUN pip install --upgrade pip && \
  pip install \
  mkdocs-git-revision-date-localized-plugin \
  mkdocs-glightbox \
  mkdocs-awesome-pages-plugin \
  pillow \
  cairosvg
