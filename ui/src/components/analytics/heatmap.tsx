import React, { useEffect, useRef } from 'react';
import Datamap from 'datamaps';
import * as d3 from 'd3';
import CanadaJson from './india.topo.json';

const ChoroplethMap = ({ data }) => {
    const mapContainerRef = useRef(null);

    useEffect(() => {
        let dataset = {};

        let onlyValues = data.map(function (obj) { return obj[1]; });
        let minValue = Math.min.apply(null, onlyValues);
        let maxValue = Math.max.apply(null, onlyValues);

        let paletteScale = d3.scaleLinear()
            .domain([minValue, maxValue])
            .range(["#EFEFFF", "#02386F"]); // blue color

        data.forEach(function (item: any) {
            let iso = item[0];
            let value = item[1];
            dataset[iso] = { numberOfThings: value, fillColor: paletteScale(value) };
        });

        let map = new Datamap({
            element: mapContainerRef.current,
            scope: 'canada',
            geographyConfig: {
                popupOnHover: true,
                highlightOnHover: true,
                borderColor: '#444',
                highlightBorderWidth: 1,
                borderWidth: 0.5,
                dataJson: CanadaJson,
                popupTemplate: function (geo, data) {
                    if (!data) { return; }
                    return [
                        '<div class="hoverinfo">',
                        '<strong>', geo.properties.name, '</strong>',
                        '<br>Count: <strong>', data.numberOfThings, '</strong>',
                        '</div>'
                    ].join('');
                }
            },
            fills: {
                HIGH: '#afafaf',
                LOW: '#123456',
                MEDIUM: 'blue',
                UNKNOWN: 'rgb(0,0,0)',
                defaultFill: '#eee'
            },
            data: dataset,
            setProjection: function (element) {
                var projection = d3.geoMercator()
                    .center([-106.3468, 68.1304]) // always in [East Latitude, North Longitude]
                    .scale(200)
                    .translate([element.offsetWidth / 2, element.offsetHeight / 2]);

                var path = d3.geoPath().projection(projection);
                return { path: path, projection: projection };
            }
        });

        // Clean up the map on component unmount
        return () => {
            map.svg.remove();
        };
    }, [data]);

    return (
        <div id="cloropleth_map" ref={mapContainerRef} style={{
            height: "100%",
            width: "100%",
        }}></div>
    );
};

export default ChoroplethMap;
