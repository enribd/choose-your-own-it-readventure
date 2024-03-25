package loader

import (
	"log"
	"sort"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
)

// LearningPaths["ref"] = LearningPath{}
var LearningPaths map[string]models.LearningPath = make(map[string]models.LearningPath)

// This structure is the same as LearningPaths but type agnostic, used in
// template data because it only allows map[string]any types
var LearningPathsTmpl map[string]any = make(map[string]any)

func loadLearningPaths(basepath string) error {
	log.Printf("load learning paths from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("learning path files %s\n", files)

	// Load the content of the files and populate the Books var
	var content []models.LearningPath
	for _, f := range files {
		content, err = loadFile[models.LearningPath](f)
		if err != nil {
			return err
		}

		for _, lp := range content {
			// skip if lp has no books or if status is coming soon
			if lp.Status == "coming-soon" || stats.Data.TotalLearningPathBooks[string(lp.Ref)] == 0 {
				stats.IncSkippedLearningPath()
			} else {
				// check for duplicates
				seenTabs := make(map[models.LearningPathTabRef]bool)
				seenTags := make(map[models.TagRef]bool)

				for _, t := range lp.Tags {
					if _, ok := seenTags[t]; ok {
						log.Fatalf("loader: %s learning path has duplicated tags: %s", lp.Ref, t)
					}
					seenTags[t] = true
				}

				// If not populated, use the default data from the tabs file
				for i, t := range lp.Tabs {
					if _, ok := seenTabs[t.Ref]; ok {
						log.Fatalf("loader: %s learning path has duplicated tabs: %s", lp.Ref, t.Ref)
					}
					seenTabs[t.Ref] = true

					if t.Data.Name == "" {
						t.Data.Name = LearningPathsTabs[t.Ref].Name
					}
					if t.Data.Icon == "" {
						t.Data.Icon = LearningPathsTabs[t.Ref].Icon
					}
					if t.Data.Desc == "" {
						t.Data.Desc = LearningPathsTabs[t.Ref].Desc
					}
					// if no description wanted set to "empty"
					if t.Data.Desc == "empty" {
						t.Data.Desc = ""
					}
					if t.Data.Order == 0 {
						t.Data.Order = LearningPathsTabs[t.Ref].Order
					}

					lp.Tabs[i] = t
				}

				LearningPaths[string(lp.Ref)] = lp
				stats.SetTotalLearningPaths(len(LearningPaths))
			}
		}
	}

	// Avoid learning paths having empty related or suggested learning paths
	purgeEmtpyRelatedLearningPaths()
	purgeEmtpySuggestedLearningPaths()

	// Sort tabs in learning path
	sortLearningPathTabs()

	// Build auxiliar template structure
	for lpRef, lp := range LearningPaths {
		LearningPathsTmpl[string(lpRef)] = lp
	}

	return nil
}

// Avoid learning paths having empty related learning paths
func purgeEmtpyRelatedLearningPaths() {
	for _, lp := range LearningPaths {
		// Remove empty related paths
		var notEmtpyLPs []models.LearningPathRef
		for _, relatedRef := range lp.Related {
			if stats.Data.TotalLearningPathBooks[string(relatedRef)] > 0 {
				notEmtpyLPs = append(notEmtpyLPs, relatedRef)
			} /* else {
							log.Printf("'%s' is an empty or a coming soon learning path, removed from '%s' related paths", relatedRef, lp.Ref)
			      } */
		}
		lp.Related = notEmtpyLPs
		LearningPaths[string(lp.Ref)] = lp
	}
}

// Avoid learning paths having empty suggested learning paths
func purgeEmtpySuggestedLearningPaths() {
	for _, lp := range LearningPaths {
		var notEmtpyLPs []models.LearningPathRef
		for _, suggestedRef := range lp.Suggested {
			// If the suggested exists in the active lps map add it to the new suggested list
			if _, ok := LearningPaths[string(suggestedRef)]; ok {
				notEmtpyLPs = append(notEmtpyLPs, suggestedRef)
			} /* else {
							log.Printf("'%s' is an empty or a coming soon learning path, removed from '%s' suggested paths", suggestedRef, lp.Ref)
			} */
		}
		lp.Suggested = notEmtpyLPs
		LearningPaths[string(lp.Ref)] = lp
	}
}

// Sort tabs by order ascendant
func sortLearningPathTabs() {
	for _, lp := range LearningPaths {
		tabs := lp.Tabs
		sort.SliceStable(tabs, func(i, j int) bool {
			if tabs[i].Data.Order == tabs[j].Data.Order {
				log.Fatalf("loader: %s learning path tabs %s and %s can't have the same order %d", lp.Ref, tabs[i].Ref, tabs[j].Ref, tabs[i].Data.Order)
			}
			return tabs[i].Data.Order < tabs[j].Data.Order
		})
	}
}
