package graphics

import (
	tps "clusterAnalysis/lib/types"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
)

func DrawPoints(clusters []tps.Cluster) {
	p := plot.New()
	p.Title.Text = "Clusters shows"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	// colors to draw clusters
	colors := plotutil.DarkColors

	for i, cluster := range clusters {
		// convert points to plot format
		pts := make(plotter.XYs, len(cluster.Points))
		for j, point := range cluster.Points {
			pts[j].X = point.X
			pts[j].Y = point.Y
		}

		// Создаем Scatter для точек кластера
		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(fmt.Errorf("ошибка создания Scatter: %v", err))
		}

		// Настраиваем стиль точек
		colorIdx := i % len(colors)
		scatter.GlyphStyle = draw.GlyphStyle{
			Color:  colors[colorIdx],
			Shape:  draw.CircleGlyph{},
			Radius: vg.Points(4),
		}

		// Добавляем точки на график
		p.Add(scatter)
		p.Legend.Add(fmt.Sprintf("Cluster %d", i+1), scatter)

		// Добавляем центроид
		centroid := plotter.XY{
			X: cluster.Centroid.X,
			Y: cluster.Centroid.Y,
		}
		centroidScatter, err := plotter.NewScatter(plotter.XYs{centroid})
		if err != nil {
			log.Fatal(fmt.Errorf("ошибка создания центроида: %v", err))
		}

		// Настраиваем стиль центроида
		centroidScatter.GlyphStyle = draw.GlyphStyle{
			Color:  color.Black,
			Shape:  draw.CrossGlyph{},
			Radius: vg.Points(8),
		}

		p.Add(centroidScatter)
	}
	// Сохраняем график в файл PNG
	if err := p.Save(600, 400, "scatter_plot.png"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Plot was saved")
}
