compname=$1
echo ${compname}

cp -r ./boilerplate ./src/components/${compname}

exts=(css html styl vue vue.ts)
for ext in ${exts[@]}; do
    mv ./src/components/${compname}/Template.component.${ext} \
       ./src/components/${compname}/${compname}.component.${ext}
done

exts=(vue vue.ts)
for ext in ${exts[@]}; do
    filepath=./src/components/${compname}/${compname}.component.${ext}
    sed -ri "s/[{]{2} ?Template ?[}]{2}/${compname}/g" ${filepath}
done
