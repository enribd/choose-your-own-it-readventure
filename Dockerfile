FROM squidfunk/mkdocs-material:9.5

RUN pip install --upgrade pip && \
  pip install \
  mkdocs-git-revision-date-localized-plugin \
  mkdocs-glightbox \
  mkdocs-awesome-pages-plugin \
  pillow \
  cairosvg
